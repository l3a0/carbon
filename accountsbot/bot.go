package accountsbot

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"gopkg.in/cenkalti/backoff.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	c "github.com/l3a0/carbon/contracts"
	m "github.com/l3a0/carbon/models"
)

// Bot represents some logic that runs in background.
type Bot interface {
	Wake(statusChannel chan int)
	Work(statusChannel chan int)
	Sleep(statusChannel chan int)
}

// AccountsBot maintains state for accounts with debt.
type AccountsBot struct {
	ethClient          *ethclient.Client
	accounts           map[string]*m.Account
	tokens             map[string]c.Token
	botsCollection     *mgo.Collection
	accountsCollection *mgo.Collection
	state              *BotState
}

// BotState represents the state of the bot.
type BotState struct {
	ID                     bson.ObjectId `bson:"_id,omitempty"`
	BotType                string
	LastWakeTime           time.Time
	LastSleepTime          time.Time
	LastBorrowBlockByToken map[string]uint64
}

// NewAccountsBot creates a new AccountsBot.
// TODO: Refactor dependencies into interfaces.
func NewAccountsBot(ethClient *ethclient.Client, botsCollection *mgo.Collection, accountsCollection *mgo.Collection) *AccountsBot {
	return &AccountsBot{
		ethClient:          ethClient,
		botsCollection:     botsCollection,
		accountsCollection: accountsCollection,
	}
}

// String returns string representation of the bot.
func (state BotState) String() string {
	value, _ := json.Marshal(&state)
	return string(value)
}

// String returns string representation of the bot.
func (bot AccountsBot) String() string {
	return fmt.Sprintf("AccountsBot{%v}", bot.state)
}

// Wake gets the bot ready for work.
func (bot *AccountsBot) Wake(statusChannel chan int) {
	log.Printf("%v waking...\n", bot)
	bot.accounts = make(map[string]*m.Account)
	// restore state for accounts bot.
	// query the db for bot with bottype == AccountsBot.
	state := BotState{}
	log.Printf("state.LastBlockParsedByTokenSymbol: %v\n", state.LastBorrowBlockByToken)
	err := bot.botsCollection.Find(bson.M{"bottype": "AccountsBot"}).One(&state)
	if err != nil {
		// log.Fatalf("Error finding record: %#v", err)
		log.Printf("Error finding record: %v\n", err)
		// create the initial bot record.
		state.ID = bson.NewObjectId()
		state.BotType = "AccountsBot"
		state.LastWakeTime = time.Now()
		log.Printf("Inserting AccountsBot: %v\n", state)
		err = bot.botsCollection.Insert(state)
		if err != nil {
			log.Fatalf("Problem inserting data: %v", err)
		}
		log.Printf("Inserted AccountsBot: %v\n", state)
	} else {
		state.LastWakeTime = time.Now()
		updateQuery := bson.M{"_id": state.ID}
		change := bson.M{"$set": bson.M{"lastwaketime": state.LastWakeTime}}
		log.Printf("Updating AccountsBot: %v\n", state)
		err = bot.botsCollection.Update(updateQuery, change)
		if err != nil {
			log.Fatalf("Error updating record: %v", err)
		}
		log.Printf("Updated AccountsBot: %v\n", state)
	}
	bot.state = &state
	log.Printf("initialized %v\n", bot)
	log.Printf("initialized accounts collection: %#v\n", bot.accounts)
	bot.tokens = make(map[string]c.Token)
	log.Printf("initialized tokens collection: %#v\n", bot.tokens)
	for tokenSymbol := range c.TokenAddresses {
		token := c.NewToken(tokenSymbol, bot.ethClient)
		if token != nil {
			name, err := token.Name(nil)
			if err != nil {
				log.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
			}
			bot.tokens[tokenSymbol] = token
			log.Printf("Initialized token %#v (%#v)\n", name, tokenSymbol)
		}
	}
	statusChannel <- 0
}

// Work puts the bot to work.
func (bot *AccountsBot) Work(status chan int) {
	log.Printf("%v working...\n", bot)
	// TODO: go routine per token contract?
	for tokenSymbol, token := range bot.tokens {
		name, err := token.Name(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve %#v token name: %#v", tokenSymbol, err)
		}
		log.Printf("Processing accounts for token %#v (%#v) @ %#v\n", name, tokenSymbol, c.TokenAddresses[tokenSymbol].Hex())
		filterOptions := &bind.FilterOpts{Start: 0, End: nil, Context: nil}
		var iter c.TokenBorrowIterator
		// An operation that may fail.
		operation := func() error {
			iter, err = token.FilterBorrowEvents(filterOptions)
			if err != nil {
				return err
			}
			return nil
		}
		err = backoff.Retry(operation, backoff.NewExponentialBackOff())
		if err != nil {
			log.Fatalf("Failed to FilterBorrowEvents for token %#v: %#v", tokenSymbol, err)
		}
		if iter != nil {
			if bot.state.LastBorrowBlockByToken == nil {
				bot.state.LastBorrowBlockByToken = make(map[string]uint64)
			}
			bot.state.LastBorrowBlockByToken[tokenSymbol] = parseAccounts(iter, tokenSymbol, bot.accounts)
		}
	}
	log.Printf("len(accounts): %#v\n", len(bot.accounts))
	// for _, account := range accounts {
	// 	fmt.Printf("Account(%#v)\n", account.address)
	// 	for tokenSymbol, tokenBorrow := range account.borrows {
	// 		fmt.Printf("Borrow: %#v (%#v)\n", tokenSymbol, tokenBorrow)
	// 	}
	// }
	status <- 0
}

// Sleep saves the bot's state and lets it rest.
func (bot *AccountsBot) Sleep(statusChannel chan int) {
	log.Printf("%v sleeping...\n", bot)
	bot.state.LastSleepTime = time.Now()
	updateQuery := bson.M{"_id": bot.state.ID}
	change := bson.M{"$set": bson.M{"lastsleeptime": bot.state.LastSleepTime, "lastborrowblockbytoken": bot.state.LastBorrowBlockByToken}}
	log.Printf("Updating AccountsBot: %v\n", bot.state)
	err := bot.botsCollection.Update(updateQuery, change)
	if err != nil {
		log.Fatalf("Error updating record: %v", err)
	}
	log.Printf("Updated AccountsBot: %v\n", bot.state)
	statusChannel <- 0
}

func parseAccounts(iter c.TokenBorrowIterator, tokenSymbol string, accounts map[string]*m.Account) uint64 {
	log.Printf("Parsing accounts...\n")
	var lastBlock uint64 = 0
	for i := 0; iter.Next(); i++ {
		borrowEvent := iter.GetEvent()
		if borrowEvent != nil {
			lastBlock = borrowEvent.GetBlockNumber()
			address := borrowEvent.GetBorrower().Hex()
			borrows := borrowEvent.GetAccountBorrows()
			account, ok := accounts[address]
			if !ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account = &m.Account{
					Address: address,
					Borrows: make(map[string]*big.Int),
				}
				account.Borrows[tokenSymbol] = borrows
				accounts[account.Address] = account
				// fmt.Printf("Added account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
			} else if ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account.Borrows[tokenSymbol] = borrows
				// fmt.Printf("Updated account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
			} else if ok && borrows.Cmp(big.NewInt(0)) < 1 {
				account.Borrows[tokenSymbol] = borrows
				// check if all token borrows for the account are 0.
				accountEmpty := true
				for _, value := range account.Borrows {
					if value.Cmp(big.NewInt(0)) == 1 {
						accountEmpty = false
					}
				}
				if accountEmpty {
					delete(accounts, address)
					log.Printf("Deleted account: %#v. Balance: %#v (%#v)\n", account.Address, borrows, tokenSymbol)
				}
			}
		}
	}
	return lastBlock
}
