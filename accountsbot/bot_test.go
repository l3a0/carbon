package accountsbot

import (
	"bytes"
	"crypto/tls"
	"log"
	"net"
	"testing"
	"time"
	"regexp"

	"github.com/l3a0/carbon/contracts"
	"gopkg.in/mgo.v2"
)

func TestAccountsBot_Wake(t *testing.T) {
	// Arrange
	tokensProvider := &contracts.MockTokenContracts{}
	database := "bao-blockchain"
	username := "bao-blockchain"
	password := ""
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
	dbSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		t.Fatalf("Can't connect to cosmos db: %#v\n", err)
	}
	// t.Logf("Conntected to cosmos db: %#v\n", dbSession)
	defer dbSession.Close()
	botsCollection := dbSession.DB(database).C("bots-test")
	accountsCollection := dbSession.DB(database).C("accounts-test")
	botsCollection.DropCollection()
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
			name: "Should create bot state, initialize accounts, and go to sleep.",
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
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Inserted AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized state for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initializing accounts for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Initialized 0 accounts for AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`AccountsBot{{"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}} sleeping...`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updating AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			re = regexp.MustCompile(`Updated AccountsBot: {"ID":"[a-f\d]{24}","BotType":"AccountsBot","LastWakeTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastSleepTime":"(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?([+-][0-2]\d:[0-5]\d)?","LastBorrowBlockByToken":null}`)
			if !re.MatchString(output) {
				t.Errorf("re.MatchString(output) = %v, want %v", re.MatchString(output), true)
			}
			output, _ = buf.ReadString('\n')
			if output != "" {
				t.Errorf("output, _ = buf.ReadString('\n') = %v, want %v", output, "")
			}
		})
	}
}