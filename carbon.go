package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"

	ab "github.com/l3a0/carbon/accountsbot"
)

func main() {
	// ethClient, err := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	ethClient, err := ethclient.Dial("https://mainnet.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("we have a connection\n")

	accountsBotID := uuid.New()
	accountsBot := ab.NewAccountsBot(accountsBotID, ethClient)
	statusChannel := make(chan int)
	go accountsBot.Work(statusChannel)

	log.Printf("statusChannel: %v\n", <-statusChannel)

	// for _, account := range accounts {
	// 	fmt.Printf("Account(%#v)\n", account.address)
	// 	for tokenSymbol, tokenBorrow := range account.borrows {
	// 		fmt.Printf("Borrow: %#v (%#v)\n", tokenSymbol, tokenBorrow)
	// 	}
	// }

	// database := "bao-blockchain"
	// username := "bao-blockchain"
	// password := ""

	// // DialInfo holds options for establishing a session with Azure Cosmos DB for MongoDB API account.
	// dialInfo := &mgo.DialInfo{
	// 	Addrs:    []string{"bao-blockchain.mongo.cosmos.azure.com:10255"}, // Get HOST + PORT
	// 	Timeout:  10 * time.Second,
	// 	Database: database,
	// 	Username: username,
	// 	Password: password,
	// 	DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
	// 		return tls.Dial("tcp", addr.String(), &tls.Config{})
	// 	},
	// }

	// Create a session which maintains a pool of socket connections
	// session, err := mgo.DialWithInfo(dialInfo)

	// if err != nil {
	// 	log.Fatalf("Can't connect, go error %#v\n", err)
	// }

	// defer session.Close()

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
