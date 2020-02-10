package accountsbot

import (
	"bytes"
	"crypto/tls"
	"log"
	"net"
	"regexp"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/l3a0/carbon/contracts"
	"gopkg.in/mgo.v2"
)

func setupDbSession(database string, username string, password string) (*mgo.Session, error) {
	// DialInfo holds options for establishing a session with Azure Cosmos DB for MongoDB API account.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"bao-blockchain.mongo.cosmos.azure.com:10255"}, // Get HOST + PORT
		Timeout:  10 * time.Second,
		Database: database,
		Username: username,
		Password: password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	// Create a session which maintains a pool of socket connections
	return mgo.DialWithInfo(dialInfo)
}

func setupDbCollections(database string, dbSession *mgo.Session) (*mgo.Collection, *mgo.Collection) {
	botsCollection := dbSession.DB(database).C("bots-test")
	accountsCollection := dbSession.DB(database).C("accounts-test")
	botsCollection.DropCollection()
	accountsCollection.DropCollection()
	return botsCollection, accountsCollection
}

func TestAccountsBot_Wake(t *testing.T) {
	// Arrange
	database := "bao-blockchain"
	username := "bao-blockchain"
	password := ""
	dbSession, err := setupDbSession(database, username, password)
	if err != nil {
		t.Fatalf("Can't connect to cosmos db: %#v\n", err)
	}
	defer dbSession.Close()
	botsCollection, accountsCollection := setupDbCollections(database, dbSession)
	borrowEvents := []contracts.TokenBorrow{
		&contracts.CUSDCBorrow{
			Borrower: common.Address{},
			BorrowAmount: common.Big0,
			AccountBorrows: common.Big1,
			TotalBorrows: common.Big0,
			Raw: types.Log{},
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
		contracts.CBATSymbol: fakeToken2,
	}
	tokensProvider := &contracts.MockTokenContracts{
		Contracts: fakeContracts,
	}
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	type fields struct {
		botsCollection     *mgo.Collection
		accountsCollection *mgo.Collection
		tokensProvider     contracts.TokensProvider
		logger             *log.Logger
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
			},
			args: args{
				statusChannel: make(chan int),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewAccountsBot(tt.fields.botsCollection, tt.fields.accountsCollection, tt.fields.tokensProvider, tt.fields.logger)
			// Act
			go bot.Wake(tt.args.statusChannel)
			status := <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			go bot.Work(tt.args.statusChannel)
			status = <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			go bot.Sleep(tt.args.statusChannel)
			status = <-tt.args.statusChannel
			if status != 0 {
				t.Errorf("status := <-tt.args.statusChannel = %v, want %v", status, 0)
			}
			// Assert
			output, _ := buf.ReadString('\n')
			if output != "AccountsBot{<nil>} waking...\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "AccountsBot{<nil>} waking...\\n")
			}
			output, _ = buf.ReadString('\n')
			if output != "Could not find existing bot state: not found\n" {
				t.Errorf("output, _ := buf.ReadString('\\n') = %v, want %v", output, "Could not find existing bot state: not found\\n")
			}
			output, _ = buf.ReadString('\n')
			re := regexp.MustCompile(`Inserting AccountsBot state: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Inserting AccountsBot state: {"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Inserted AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Inserted AccountsBot: {"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized state for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Initialized state for AccountsBot{{"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{}}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initializing accounts for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, `Initializing accounts for AccountsBot{{"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized 0 accounts for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initialized 0 accounts for AccountsBot{{"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}} working...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Initialized 0 accounts for AccountsBot{{"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":null}}`)
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
			re = regexp.MustCompile(`Upserting account: &{ObjectIdHex\("[a-f\d]{24}"\) 0x0000000000000000000000000000000000000000 map\[CBAT:1 CUSDC:1\]}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Upserting account: &{ObjectIdHex("*") 0x0000000000000000000000000000000000000000 map[CBAT:1 CUSDC:1]}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Upserted account: &{ObjectIdHex\("[a-f\d]{24}"\) 0x0000000000000000000000000000000000000000 map\[CBAT:1 CUSDC:1\]}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Upserted account: &{ObjectIdHex("*") 0x0000000000000000000000000000000000000000 map[CBAT:1 CUSDC:1]}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{}}} sleeping...`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `AccountsBot{{"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{}}} sleeping...`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updating AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Updating AccountsBot: {"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{}}`)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updated AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":{}}`)
			if !re.MatchString(output) {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, `Updated AccountsBot: {"ID":"*","BotType":"AccountsBot","LastWakeTime":"*","LastSleepTime":"*","LastBorrowBlockByToken":{}}`)
			}
			output, _ = buf.ReadString('\n')
			if output != "" {
				t.Errorf("output, _ = buf.ReadString('\\n') = %v, want %v", output, "")
			}
			botsCollection.DropCollection()
			accountsCollection.DropCollection()
		})
	}
}