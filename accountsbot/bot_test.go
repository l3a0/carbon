package accountsbot

import (
	"bytes"
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/l3a0/carbon/contracts"
	"github.com/l3a0/carbon/models"
)

func TestAccountsBot_Wake(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	// logger := log.New(os.Stderr, "", 0)
	cosmosClient := models.NewCosmosService(
		// log.New(os.Stderr, "CosmosClient | ", 0),
		logger,
		models.CosmosConfiguration{
			SubscriptionID:    "6951e94d-0947-4c5e-b865-f864609da246",
			CloudName:         "AzurePublicCloud",
			ResourceGroupName: "crmbts-devmachines-bao-blockchain-460259",
			AccountName:       "bao-blockchain",
		})
	cosmosClient.Connect()
	ctx := context.Background()
	session, err := cosmosClient.GetSession(ctx)
	if err != nil {
		t.Errorf("cannot get mongoDB session: %v", err)
	}
	defer session.Close()
	documentDbCollectionFactory := models.NewCosmosCollectionFactory(
		// log.New(os.Stderr, "DocumentDbCollectionFactory | ", 0),
		logger,
		cosmosClient,
		session)
	botsCollectionName := "mock-bots"
	accountsCollectionName := "mock-accounts"
	botsCollection, err := documentDbCollectionFactory.CreateCollection(ctx, botsCollectionName)
	if err != nil {
		t.Errorf("Could not create collection: %v", err)
	}
	accountsCollection, err := documentDbCollectionFactory.CreateCollection(ctx, accountsCollectionName)
	if err != nil {
		t.Errorf("Could not create collection: %v", err)
	}
	// Arrange
	borrowEvents := []contracts.TokenBorrow{
		&contracts.CUSDCBorrow{
			Borrower:       common.Address{},
			BorrowAmount:   common.Big0,
			AccountBorrows: common.Big1,
			TotalBorrows:   common.Big0,
			Raw:            types.Log{},
		},
	}
	fakeTokenBorrowIterator := &contracts.MockTokenBorrowIterator{
		BorrowEvents: borrowEvents,
	}
	fakeTokenBorrowIterator2 := &contracts.MockTokenBorrowIterator{
		BorrowEvents: borrowEvents,
	}
	fakeToken := &contracts.MockToken{
		TokenBorrowIterator: fakeTokenBorrowIterator,
	}
	fakeToken2 := &contracts.MockToken{
		TokenBorrowIterator: fakeTokenBorrowIterator2,
	}
	fakeContracts := map[string]contracts.Token{
		contracts.CUSDCSymbol: fakeToken,
		contracts.CBATSymbol:  fakeToken2,
	}
	tokensProvider := &contracts.MockTokenContracts{
		Contracts: fakeContracts,
	}
	accountsService := models.NewCosmosAccountsService(
		// log.New(os.Stderr, "CosmosAccountsService | ", 0),
		logger,
		documentDbCollectionFactory,
		accountsCollectionName)
	botsService := models.NewCosmosBotsService(
		// log.New(os.Stderr, "CosmosBotsService | ", 0),
		logger,
		documentDbCollectionFactory,
		botsCollectionName)
	type fields struct {
		botsCollection     models.Collection
		accountsCollection models.Collection
		tokensProvider     contracts.TokensProvider
		logger             *log.Logger
		accountsService    models.AccountsService
		botsService        models.BotsService
	}
	type args struct {
		statusChannel chan int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Should wake, work, and sleep.",
			fields: fields{
				botsCollection:     botsCollection,
				accountsCollection: accountsCollection,
				tokensProvider:     tokensProvider,
				logger:             logger,
				accountsService:    accountsService,
				botsService:        botsService,
			},
			args: args{
				statusChannel: make(chan int),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewAccountsBot(
				tt.fields.tokensProvider,
				tt.fields.logger,
				tt.fields.accountsService,
				tt.fields.botsService)
			// Act
			go bot.Wake(ctx, tt.args.statusChannel)
			status := <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			go bot.Work(ctx, tt.args.statusChannel)
			status = <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			go bot.Sleep(ctx, tt.args.statusChannel)
			status = <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			go bot.Wake(ctx, tt.args.statusChannel)
			status = <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			// Assert
			output, _ := buf.ReadString('\n')
			if output != "Collection not found: mock-bots.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Collection not found: mock-bots.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Creating collection: mock-bots.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Creating collection: mock-bots.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Created collection: mock-bots.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Created collection: mock-bots.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Collection not found: mock-accounts.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Collection not found: mock-accounts.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Creating collection: mock-accounts.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Creating collection: mock-accounts.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Created collection: mock-accounts.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Created collection: mock-accounts.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "AccountsBot{<nil>} waking...\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "AccountsBot{<nil>} waking...\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Collection found: mock-bots.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Collection found: mock-bots.\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Could not find existing bot state: not found\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Could not find existing bot state: not found\\n")
			}
			output, _ = buf.ReadString('\n')
			re := regexp.MustCompile(`Inserting AccountsBot state: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Inserting AccountsBot state: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Creating Bot State: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Creating Bot State: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Created Bot State: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Created Bot State: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Inserted AccountsBot: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Inserted AccountsBot: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized state for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Initialized state for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{}}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initializing accounts for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Initializing accounts for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}}`)
			}
			output, _ = buf.ReadString('\n')
			if output != "Collection found: mock-accounts.\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Collection found: mock-accounts.\\n")
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized 0 accounts for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initialized 0 accounts for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}} working...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}} working...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Processing accounts for token MockToken \([A-Z]+\) at block # 0 @ 0x0000000000000000000000000000000000000000`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Processing accounts for token MockToken (*) at block # 0 @ 0x0000000000000000000000000000000000000000`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Parsing accounts...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Parsing accounts...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Added account: "0x0000000000000000000000000000000000000000". Borrowed 1 \("[A-Z]+"\)`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Added account: "0x0000000000000000000000000000000000000000". Borrowed 1 ("*")`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Processing accounts for token MockToken \([A-Z]+\) at block # 0 @ 0x0000000000000000000000000000000000000000`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Processing accounts for token MockToken (*) at block # 0 @ 0x0000000000000000000000000000000000000000`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Parsing accounts...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Parsing accounts...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updated account: "0x0000000000000000000000000000000000000000". Borrowed 1 \("[A-Z]+"\)`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Updated account: "0x0000000000000000000000000000000000000000". Borrowed 1 ("*")`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Upserting account: &{ObjectIdHex\("[a-f\d]{24}"\) 0x0000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000 map\[CBAT:1 CUSDC:1\]}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Upserting account: &{ObjectIdHex("*") 0x0000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000 map[CBAT:1 CUSDC:1]}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Upserted account: &{ObjectIdHex\("[a-f\d]{24}"\) 0x0000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000 map\[CBAT:1 CUSDC:1\]}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Upserted account: &{ObjectIdHex("*") 0x0000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000 map[CBAT:1 CUSDC:1]}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`numberOfModifiedAccounts: 1`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `numberOfModifiedAccounts: 1`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`numberOfAccounts: 1`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `numberOfAccounts: 1`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}} sleeping...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}} sleeping...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updating AccountsBot: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Updating AccountsBot: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updated AccountsBot: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Updated AccountsBot: {"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}} waking...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}} waking...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updating AccountsBot: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updated AccountsBot: {"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized state for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initialized state for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initializing accounts for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initializing accounts for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized 1 accounts for AccountsBot{{"ShardKey":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initialized 1 accounts for AccountsBot{{"ShardKey":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{"CBAT":0,"CUSDC":0}}}`)
			}
			err = cosmosClient.DeleteSQLContainer(ctx, botsCollectionName)
			if err != nil {
				t.Errorf("Failed to delete container: %v", err)
			}
			err = cosmosClient.DeleteSQLContainer(ctx, accountsCollectionName)
			if err != nil {
				t.Errorf("Failed to delete container: %v", err)
			}
		})
	}
}
