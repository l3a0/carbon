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
		log.New(os.Stderr, "CosmosClient | ", log.LstdFlags),
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
		log.New(os.Stderr, "DocumentDbCollectionFactory | ", log.LstdFlags),
		cosmosClient,
		session)
	var accountsBot accountsbot.Bot
	tokenContracts, err := contracts.NewTokenContracts(
		ethClient,
		log.New(os.Stderr, "TokenFactory | ", log.LstdFlags),
		contracts.NewToken)
	if err != nil {
		log.Panic(err)
	}
	botsCollectionName := "bots"
	accountsCollectionName := "accounts"
	accountsService := models.NewCosmosAccountsService(
		log.New(os.Stderr, "CosmosAccountsService | ", log.LstdFlags),
		documentDbCollectionFactory,
		accountsCollectionName)
	botsService := models.NewCosmosBotsService(
		log.New(os.Stderr, "DocumentDbCollectionFactory | ", log.LstdFlags),
		documentDbCollectionFactory,
		botsCollectionName)
	comptrollerService := models.NewComptrollerService(
		log.New(os.Stderr, "ComptrollerService | ", log.LstdFlags),
		ethClient)
	accountsBot = accountsbot.NewAccountsBot(
		tokenContracts,
		log.New(os.Stderr, "AccountsBot | ", log.LstdFlags),
		accountsService,
		botsService,
		comptrollerService)
	status := make(chan int)
	go accountsBot.Wake(ctx, status)
	log.Printf("Wake status: %v\n", <-status)
	go accountsBot.Work(ctx, status)
	log.Printf("Work status: %v\n", <-status)
	go accountsBot.Sleep(ctx, status)
	log.Printf("Sleep status: %v\n", <-status)
}
