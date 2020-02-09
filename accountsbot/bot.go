package accountsbot

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"gopkg.in/cenkalti/backoff.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/l3a0/carbon/contracts"
	"github.com/l3a0/carbon/models"
)

// Bot represents some logic that runs in background.
type Bot interface {
	Wake(statusChannel chan int)
	Work(statusChannel chan int)
	Sleep(statusChannel chan int)
}

// AccountsBot maintains state for accounts with debt.
type AccountsBot struct {
	accounts           map[string]*models.Account
	tokens             map[string]contracts.Token
	tokenAddresses     map[string]common.Address
	botsCollection     *mgo.Collection
	accountsCollection *mgo.Collection
	state              *BotState
	logger             *log.Logger
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
func NewAccountsBot(botsCollection *mgo.Collection, accountsCollection *mgo.Collection, tokensProvider contracts.TokensProvider, logger *log.Logger) *AccountsBot {
	return &AccountsBot{
		botsCollection:     botsCollection,
		accountsCollection: accountsCollection,
		tokens:             tokensProvider.GetTokens(),
		tokenAddresses:     tokensProvider.GetAddresses(),
		logger:             logger,
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

func (bot *AccountsBot) insertState(state *BotState) {
	// create the initial bot record.
	state.ID = bson.NewObjectId()
	state.BotType = "AccountsBot"
	state.LastWakeTime = time.Now()
	// An operation that may fail.
	operation := func() error {
		bot.logger.Printf("Inserting AccountsBot state: %v\n", state)
		err := bot.botsCollection.Insert(state)
		if err != nil {
			bot.logger.Printf("Problem inserting data: %T %v %v", err, err, err.(*mgo.LastError))
			return err
		}
		return nil
	}
	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		bot.logger.Fatalf("Error updating record: %T %v", err, err)
	}
	bot.logger.Printf("Inserted AccountsBot: %v\n", state)
}

func (bot *AccountsBot) initializeState() {
	// restore state for accounts bot.
	// query the db for bot with bottype == AccountsBot.
	state := BotState{}
	err := bot.botsCollection.Find(bson.M{"bottype": "AccountsBot"}).One(&state)
	if err != nil {
		bot.logger.Printf("Could not find existing bot state: %v\n", err)
		bot.insertState(&state)
	} else {
		state.LastWakeTime = time.Now()
		updateQuery := bson.M{"_id": state.ID}
		change := bson.M{"$set": bson.M{"lastwaketime": state.LastWakeTime}}
		// An operation that may fail.
		operation := func() error {
			bot.logger.Printf("Updating AccountsBot: %v\n", state)
			err = bot.botsCollection.Update(updateQuery, change)
			if err != nil {
				bot.logger.Printf("Error updating record: %T %v %v", err, err, err.(*mgo.LastError))
				return err
			}
			return nil
		}
		err = backoff.Retry(operation, backoff.NewExponentialBackOff())
		if err != nil {
			bot.logger.Fatalf("Error updating record: %T %v %v", err, err, err.(*mgo.LastError))
		}
		bot.logger.Printf("Updated AccountsBot: %v\n", state)
	}
	bot.state = &state
	bot.logger.Printf("Initialized state for %v\n", bot)
}

func (bot *AccountsBot) initializeAccounts() {
	// restore accounts from db.
	bot.logger.Printf("Initializing accounts for %v.\n", bot)
	bot.accounts = make(map[string]*models.Account)
	var accounts []*models.Account
	err := bot.accountsCollection.Find(nil).All(&accounts)
	if err != nil {
		bot.logger.Fatalf("Error finding accounts: %v\n", err)
	}
	for _, account := range accounts {
		bot.accounts[account.Address] = account
	}
	bot.logger.Printf("Initialized %v accounts for %v.\n", len(bot.accounts), bot)
}

// Wake gets the bot ready for work.
func (bot *AccountsBot) Wake(statusChannel chan int) {
	bot.logger.Printf("%v waking...\n", bot)
	bot.initializeState()
	bot.initializeAccounts()
	statusChannel <- 0
}

func (bot *AccountsBot) filterBorrowEvents(tokenSymbol string, tokenName string, token contracts.Token) contracts.TokenBorrowIterator {
	log.Printf("Processing accounts for token %v (%v) at block # %v @ %v\n", tokenName, tokenSymbol, bot.state.LastBorrowBlockByToken[tokenSymbol], bot.tokenAddresses[tokenSymbol].Hex())
	// +1 => exclude the last borrow block.
	filterOptions := &bind.FilterOpts{Start: bot.state.LastBorrowBlockByToken[tokenSymbol], End: nil, Context: nil}
	var iter contracts.TokenBorrowIterator
	var err error
	// An operation that may fail.
	operation := func() error {
		iter, err = token.FilterBorrowEvents(filterOptions)
		if err != nil {
			log.Printf("Failed to FilterBorrowEvents for token %v: %v", tokenSymbol, err)
			return err
		}
		return nil
	}
	err = backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		log.Fatalf("Failed to FilterBorrowEvents for token %v: %v", tokenSymbol, err)
	}
	return iter
}

// Work puts the bot to work.
func (bot *AccountsBot) Work(status chan int) {
	bot.logger.Printf("%v working...\n", bot)
	// TODO: go routine per token contract?
	modifiedAccounts := []*models.Account{}
	for tokenSymbol, token := range bot.tokens {
		tokenName, err := token.Name(nil)
		if err != nil {
			bot.logger.Fatalf("Failed to retrieve %v token name: %v", tokenSymbol, err)
		}
		iter := bot.filterBorrowEvents(tokenSymbol, tokenName, token)
		if iter != nil {
			if bot.state.LastBorrowBlockByToken == nil {
				bot.state.LastBorrowBlockByToken = make(map[string]uint64)
			}
			block, accounts := bot.parseAccounts(iter, tokenSymbol)
			modifiedAccounts = append(modifiedAccounts, accounts...)
			if block > bot.state.LastBorrowBlockByToken[tokenSymbol] {
				bot.state.LastBorrowBlockByToken[tokenSymbol] = block
			}
		}
	}
	numberOfModifiedAccounts := len(modifiedAccounts)
	bot.logger.Printf("numberOfModifiedAccounts: %v\n", numberOfModifiedAccounts)
	numberOfAccounts := len(bot.accounts)
	bot.logger.Printf("numberOfAccounts: %v\n", numberOfAccounts)
	if numberOfAccounts > 0 {
		for _, account := range modifiedAccounts {
			operation := func() error {
				fmt.Printf("Upserting account: %v\n", account)
				_, err := bot.accountsCollection.Upsert(bson.M{"_id": account.ID}, account)
				if err != nil {
					bot.logger.Printf("Problem upserting data: %T %v %v", err, err, err.(*mgo.LastError))
					return err
				}
				bot.logger.Printf("Upserted account: %v\n", account)
				return nil
			}
			err := backoff.Retry(operation, backoff.NewExponentialBackOff())
			if err != nil {
				bot.logger.Fatalf("Problem upserting data: %T %v %v", err, err, err.(*mgo.LastError))
			}
		}
	}
	status <- 0
}

// Sleep saves the bot's state and lets it rest.
func (bot *AccountsBot) Sleep(statusChannel chan int) {
	bot.logger.Printf("%v sleeping...\n", bot)
	bot.state.LastSleepTime = time.Now()
	updateQuery := bson.M{"_id": bot.state.ID}
	change := bson.M{"$set": bson.M{"lastsleeptime": bot.state.LastSleepTime, "lastborrowblockbytoken": bot.state.LastBorrowBlockByToken}}
	bot.logger.Printf("Updating AccountsBot: %v\n", bot.state)
	err := bot.botsCollection.Update(updateQuery, change)
	if err != nil {
		bot.logger.Fatalf("Error updating record: %T %v", err, err)
	}
	bot.logger.Printf("Updated AccountsBot: %v\n", bot.state)
	statusChannel <- 0
}

func (bot *AccountsBot) parseAccounts(iter contracts.TokenBorrowIterator, tokenSymbol string) (uint64, []*models.Account) {
	bot.logger.Printf("Parsing accounts...\n")
	var lastBlock uint64 = 0
	modifiedAccounts := map[string]*models.Account{}
	for i := 0; iter.Next(); i++ {
		borrowEvent := iter.GetEvent()
		if borrowEvent != nil {
			lastBlock = borrowEvent.GetBlockNumber()
			address := borrowEvent.GetBorrower().Hex()
			borrows := borrowEvent.GetAccountBorrows()
			account, ok := bot.accounts[address]
			if !ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account = &models.Account{
					ID:      bson.NewObjectId(),
					Address: address,
					Borrows: make(map[string]*big.Int),
				}
				account.Borrows[tokenSymbol] = borrows
				bot.accounts[account.Address] = account
				modifiedAccounts[account.Address] = account
				// fmt.Printf("Added account: %#v. Borrowed %#v (%#v)\n", account.address, account.borrows[tokenSymbol], tokenSymbol)
			} else if ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account.Borrows[tokenSymbol] = borrows
				modifiedAccounts[account.Address] = account
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
					delete(bot.accounts, address)
					// TODO: enable deleting from cosmos db.
					// modifiedAccounts = append(modifiedAccounts, account)
					bot.logger.Printf("Deleted account: %#v. Balance: %#v (%#v)\n", account.Address, borrows, tokenSymbol)
				}
			}
		}
	}
	result := []*models.Account{}
	for _, account := range modifiedAccounts {
		result = append(result, account)
	}
	return lastBlock, result
}
