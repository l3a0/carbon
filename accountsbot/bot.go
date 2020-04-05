package accountsbot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/globalsign/mgo/bson"
	"gopkg.in/cenkalti/backoff.v2"

	"github.com/l3a0/carbon/contracts"
	"github.com/l3a0/carbon/models"
)

// Bot represents some logic that runs in background.
type Bot interface {
	Wake(ctx context.Context, statusChannel chan int)
	Work(ctx context.Context, statusChannel chan int)
	Sleep(ctx context.Context, statusChannel chan int)
}

// AccountsBot maintains state for accounts with debt.
type AccountsBot struct {
	accounts           map[string]*models.Account
	tokens             map[string]contracts.Token
	tokenAddresses     map[string]common.Address
	botsService        models.BotsService
	accountsService    models.AccountsService
	comptrollerService models.ComptrollerService
	state              *BotState
	logger             *log.Logger
}

// BotState represents the state of the bot.
type BotState struct {
	ShardKey               string
	BotType                string
	LastWakeTime           time.Time
	LastSleepTime          time.Time
	LastBorrowBlockByToken map[string]uint64
}

// GetShardKey returns the shard key.
func (state *BotState) GetShardKey() string {
	return state.ShardKey
}

// NewAccountsBot creates a new AccountsBot.
func NewAccountsBot(
	tokensProvider contracts.TokensProvider,
	logger *log.Logger,
	accountsService models.AccountsService,
	botsService models.BotsService,
	comptrollerService models.ComptrollerService) Bot {
	return &AccountsBot{
		botsService:        botsService,
		accountsService:    accountsService,
		comptrollerService: comptrollerService,
		tokens:             tokensProvider.GetTokens(),
		tokenAddresses:     tokensProvider.GetAddresses(),
		logger:             logger,
	}
}

// Wake gets the bot ready for work.
func (bot *AccountsBot) Wake(ctx context.Context, statusChannel chan int) {
	bot.logger.Printf("%v waking...\n", bot)
	bot.initializeState(ctx)
	bot.initializeAccounts(ctx)
	statusChannel <- 0
}

// Work puts the bot to work.
func (bot *AccountsBot) Work(ctx context.Context, status chan int) {
	bot.logger.Printf("%v working...\n", bot)
	// TODO: go routine per token contract?
	modifiedAccounts := map[string]*models.Account{}
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
			bot.state.LastBorrowBlockByToken[tokenSymbol] = bot.parseAccountBorrowBalancesFromBorrowEvents(iter, tokenSymbol, modifiedAccounts)
		}
	}
	numberOfModifiedAccounts := len(modifiedAccounts)
	numberOfAccounts := len(bot.accounts)
	// Assert the invariant: number of modified accounts is not greater than the total number of accounts.
	if numberOfModifiedAccounts > numberOfAccounts {
		bot.logger.Panicf("numberOfModifiedAccounts > numberOfAccounts.\n")
	}
	bot.logger.Printf("numberOfModifiedAccounts: %v\n", numberOfModifiedAccounts)
	bot.logger.Printf("numberOfAccounts: %v\n", numberOfAccounts)
	if numberOfModifiedAccounts > 0 {
		// TODO: go routine per account w/ bounded parallelism?
		for _, account := range modifiedAccounts {
			// TODO: move backoff logic into accounts service.
			operation := func() error {
				bot.logger.Printf("Upserting account: %v\n", account)
				err := bot.accountsService.UpsertAccount(ctx, account)
				if err != nil {
					bot.logger.Printf("Problem upserting data: %v", err)
					return err
				}
				bot.logger.Printf("Upserted account: %v\n", account)
				return nil
			}
			err := backoff.Retry(operation, backoff.NewExponentialBackOff())
			if err != nil {
				bot.logger.Panicf("Problem upserting data: %v", err)
			}
		}
	}
	for _, account := range bot.accounts {
		totalBorrows := big.NewInt(0)
		for _, tokenBorrows := range account.Borrows {
			totalBorrows = totalBorrows.Add(totalBorrows, tokenBorrows)
		}
		// Assert the invariant: no account has zero total borrows across all tokens.
		if totalBorrows.Cmp(common.Big0) <= 0 {
			bot.logger.Panicf("Account %v has totalBorrows = %v.\n", account, totalBorrows)
		}
		bot.liquidateAccount(account)
	}
	status <- 0
}

func (bot *AccountsBot) liquidateAccount(account *models.Account) {
	bot.getAccountLiquidity(account)
	if account.Liquidity.Cmp(common.Big0) == 0 && account.Shortfall.Cmp(common.Big0) == 0 {

	}
}

func (bot *AccountsBot) getAccountLiquidity(account *models.Account) {
	bot.logger.Printf("Getting liquidity for account: %v", account)
	errorCode, liquidity, shortfall, err := bot.comptrollerService.GetAccountLiquidity(nil, common.HexToAddress(account.Address))
	if err != nil || errorCode.Cmp(big.NewInt(0)) > 0 {
		bot.logger.Panicf("Problem getting account liquidity: %v, errorCode = %v", err, errorCode)
	}
	account.Liquidity = liquidity
	account.Shortfall = shortfall
	bot.logger.Printf("Liquidity for account: %v", account)
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

func (bot *AccountsBot) insertState(ctx context.Context, state *BotState) {
	// create the initial bot record.
	state.ShardKey = bson.NewObjectId().Hex()
	state.BotType = "AccountsBot"
	state.LastWakeTime = time.Now()
	// An operation that may fail.
	operation := func() error {
		bot.logger.Printf("Inserting AccountsBot state: %v\n", state)
		err := bot.botsService.CreateBotState(ctx, state)
		if err != nil {
			bot.logger.Printf("Problem inserting data: %v", err)
			return err
		}
		return nil
	}
	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		bot.logger.Panicf("Problem inserting data: %v", err)
	}
	bot.logger.Printf("Inserted AccountsBot: %v\n", state)
}

func (bot *AccountsBot) initializeState(ctx context.Context) {
	// restore state for accounts bot.
	// query the db for bot with bottype == AccountsBot.
	state := &BotState{}
	err := bot.botsService.GetBotState(ctx, state)
	if err != nil {
		bot.logger.Printf("Could not find existing bot state: %v\n", err)
		bot.insertState(ctx, state)
	} else {
		state.LastWakeTime = time.Now()
		updateQuery := bson.M{"shardkey": state.ShardKey}
		change := bson.M{"$set": bson.M{"lastwaketime": state.LastWakeTime}}
		// An operation that may fail.
		operation := func() error {
			bot.logger.Printf("Updating AccountsBot: %v\n", state)
			err = bot.botsService.UpdateBotState(ctx, updateQuery, change)
			if err != nil {
				bot.logger.Printf("Error updating record: %v", err)
				return err
			}
			return nil
		}
		err = backoff.Retry(operation, backoff.NewExponentialBackOff())
		if err != nil {
			bot.logger.Panicf("Error updating record: %v", err)
		}
		bot.logger.Printf("Updated AccountsBot: %v\n", state)
	}
	bot.state = state
	bot.logger.Printf("Initialized state for %v\n", bot)
}

func (bot *AccountsBot) initializeAccounts(ctx context.Context) {
	// restore accounts from db.
	bot.logger.Printf("Initializing accounts for %v.\n", bot)
	bot.accounts = make(map[string]*models.Account)
	accounts := []*models.Account{}
	err := bot.accountsService.GetAccounts(ctx, &accounts)
	if err != nil {
		bot.logger.Fatalf("Error finding accounts: %v\n", err)
	}
	for _, account := range accounts {
		bot.accounts[account.Address] = account
	}
	bot.logger.Printf("Initialized %v accounts for %v.\n", len(bot.accounts), bot)
}

func (bot *AccountsBot) filterBorrowEvents(tokenSymbol string, tokenName string, token contracts.Token) contracts.TokenBorrowIterator {
	bot.logger.Printf("Processing accounts for token %v (%v) at block # %v @ %v\n", tokenName, tokenSymbol, bot.state.LastBorrowBlockByToken[tokenSymbol], bot.tokenAddresses[tokenSymbol].Hex())
	// alternatively, +1 => exclude the last borrow block.
	filterOptions := &bind.FilterOpts{Start: bot.state.LastBorrowBlockByToken[tokenSymbol], End: nil, Context: nil}
	var iter contracts.TokenBorrowIterator
	var err error
	// An operation that may fail.
	operation := func() error {
		iter, err = token.FilterBorrowEvents(filterOptions)
		if err != nil {
			bot.logger.Printf("Failed to FilterBorrowEvents for token %v: %v", tokenSymbol, err)
			return err
		}
		return nil
	}
	err = backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		bot.logger.Fatalf("Failed to FilterBorrowEvents for token %v: %v", tokenSymbol, err)
	}
	return iter
}

// Sleep saves the bot's state and lets it rest.
func (bot *AccountsBot) Sleep(ctx context.Context, statusChannel chan int) {
	bot.logger.Printf("%v sleeping...\n", bot)
	bot.state.LastSleepTime = time.Now()
	updateQuery := bson.M{"shardkey": bot.state.ShardKey}
	change := bson.M{"$set": bson.M{"lastsleeptime": bot.state.LastSleepTime, "lastborrowblockbytoken": bot.state.LastBorrowBlockByToken}}
	bot.logger.Printf("Updating AccountsBot: %v\n", bot.state)
	operation := func() error {
		err := bot.botsService.UpdateBotState(ctx, updateQuery, change)
		if err != nil {
			bot.logger.Printf("Error updating record: %T %v", err, err)
			return err
		}
		bot.logger.Printf("Updated AccountsBot: %v\n", bot.state)
		return nil
	}
	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		bot.logger.Panicf("Error updating record: %T %v", err, err)
	}
	statusChannel <- 0
}

func (bot *AccountsBot) parseAccountBorrowBalancesFromBorrowEvents(iter contracts.TokenBorrowIterator, tokenSymbol string, modifiedAccounts map[string]*models.Account) uint64 {
	bot.logger.Printf("Parsing accounts...\n")
	var lastBlock uint64 = 0
	for i := 0; iter.Next(); i++ {
		borrowEvent := iter.GetEvent()
		if borrowEvent != nil {
			lastBlock = borrowEvent.GetBlockNumber()
			addressHex := borrowEvent.GetBorrower().Hex()
			borrows := borrowEvent.GetAccountBorrows()
			account, ok := bot.accounts[addressHex]
			if !ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account = &models.Account{
					ID:       bson.NewObjectId(),
					ShardKey: addressHex,
					Address:  addressHex,
					Borrows:  make(map[string]*big.Int),
				}
				account.Borrows[tokenSymbol] = borrows
				bot.accounts[account.Address] = account
				modifiedAccounts[account.Address] = account
				bot.logger.Printf("Added account: %#v. Borrowed %#v (%#v)\n", account.Address, account.Borrows[tokenSymbol], tokenSymbol)
			} else if ok && borrows.Cmp(big.NewInt(0)) == 1 {
				account.Borrows[tokenSymbol] = borrows
				modifiedAccounts[account.Address] = account
				bot.logger.Printf("Updated account: %#v. Borrowed %#v (%#v)\n", account.Address, account.Borrows[tokenSymbol], tokenSymbol)
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
					delete(bot.accounts, addressHex)
					// TODO: enable deleting from cosmos db.
					modifiedAccounts[account.Address] = nil
					bot.logger.Printf("Deleted account: %#v. Balance: %#v (%#v)\n", account.Address, borrows, tokenSymbol)
				}
			}
		}
	}
	return lastBlock
}
