package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/l3a0/carbon/accountsbot"
	"github.com/l3a0/carbon/contracts"
	"github.com/l3a0/carbon/models"
)

func main() {
	ethClient, err := ethclient.Dial("/home/l3a0/.ethereum/geth.ipc")
	// ethClient, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("we have a connection\n")
	cosmosClient := models.NewCosmosService(
		log.New(os.Stderr, "CosmosClient | ", 0),
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
		log.Panicf("cannot get mongoDB session: %v", err)
	}
	defer session.Close()
	documentDbCollectionFactory := models.NewCosmosCollectionFactory(
		log.New(os.Stderr, "DocumentDbCollectionFactory | ", 0),
		cosmosClient,
		session)
	botsCollectionName := "bots"
	botsCollection, err := documentDbCollectionFactory.CreateCollection(ctx, botsCollectionName)
	if err != nil {
		log.Panicf("Could not create collection: %v", err)
	}
	accountsCollection, err := documentDbCollectionFactory.CreateCollection(ctx, "accounts")
	if err != nil {
		log.Panicf("Could not create collection: %v", err)
	}
	logger := log.New(os.Stderr, "", log.LstdFlags)
	var accountsBot accountsbot.Bot
	tokenContracts, err := contracts.NewTokenContracts(ethClient, logger, contracts.NewToken)
	if err != nil {
		log.Panic(err)
	}
	var accountsService models.AccountsService
	botsService := models.NewCosmosBotsService(
		log.New(os.Stderr, "DocumentDbCollectionFactory | ", 0),
		documentDbCollectionFactory,
		botsCollectionName)
	accountsBot = accountsbot.NewAccountsBot(botsCollection, accountsCollection, tokenContracts, logger, accountsService, botsService)
	status := make(chan int)
	go accountsBot.Wake(ctx, status)
	log.Printf("Wake status: %v\n", <-status)
	go accountsBot.Work(ctx, status)
	log.Printf("Work status: %v\n", <-status)
	go accountsBot.Sleep(ctx, status)
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
