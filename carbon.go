package main

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	ab "github.com/l3a0/carbon/accountsbot"
	c "github.com/l3a0/carbon/contracts"

	"gopkg.in/mgo.v2"
)

func main() {
	ethClient, err := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	// ethClient, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("we have a connection\n")
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
		log.Fatalf("Can't connect to cosmos db: %#v\n", err)
	}
	// log.Printf("Conntected to cosmos db: %#v\n", session)
	defer dbSession.Close()
	botsCollection := dbSession.DB("bao-blockchain").C("bots")
	accountsCollection := dbSession.DB("bao-blockchain").C("accounts")
	tokenContracts, err := c.NewTokenContracts(ethClient, c.NewToken)
	if err != nil {
		log.Fatal(err)
	}
	var accountsBot ab.Bot
	accountsBot = ab.NewAccountsBot(botsCollection, accountsCollection, tokenContracts)
	status := make(chan int)
	go accountsBot.Wake(status)
	log.Printf("Wake status: %v\n", <-status)
	go accountsBot.Work(status)
	log.Printf("Work status: %v\n", <-status)
	go accountsBot.Sleep(status)
	log.Printf("Sleep status: %v\n", <-status)

	// comptrollerAddress := common.HexToAddress("0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b")
	// comptroller, err := NewComptroller(comptrollerAddress, ethClient)

	// if err != nil {
	//     log.Fatalf("Failed to load comptroller: %#v\n", err)
	// }

	// fmt.Printf("comptroller: %#v\n", comptroller)

	// i := 0

	// for key, element := range accounts {
	//     if i > 5 {
	//         break
	//     }

	//     fmt.Println("Key:", key, "=>", "Element:", element)

	//     error, liquidity, shortfall, err := comptroller.GetAccountLiquidity(nil, common.HexToAddress(element.address))

	//     if err != nil {
	//         log.Fatalf("Failed to GetAccountLiquidity: %#v\n", err)
	//     }

	//     fmt.Printf("error: %#v liquidity: %#v shortfall: %#v\n", error, liquidity, shortfall)
	//     i++
	// }
	// }
}
