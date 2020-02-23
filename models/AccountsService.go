package models

import (
	"context"
	"log"

	"github.com/globalsign/mgo/bson"
)

// AccountsService is responsible for CRUD on account data.
type AccountsService interface {
	GetAccounts(ctx context.Context, result interface{}) error
	UpsertAccount(ctx context.Context, account *Account) error
}

// CosmosAccountsService works against Cosmos DB SQL Core.
type CosmosAccountsService struct {
	logger                 *log.Logger
	collectionFactory      CollectionFactory
	accountsCollection     Collection
	accountsCollectionName string
}

// NewCosmosAccountsService creats a new AccountsService.
func NewCosmosAccountsService(logger *log.Logger, collectionFactory CollectionFactory, accountsCollectionName string) AccountsService {
	return &CosmosAccountsService{
		logger:                 logger,
		collectionFactory:      collectionFactory,
		accountsCollectionName: accountsCollectionName,
	}
}

// GetAccounts returns all accounts.
func (service *CosmosAccountsService) GetAccounts(ctx context.Context, result interface{}) error {
	if service.accountsCollection == nil {
		collection, err := service.collectionFactory.CreateCollection(ctx, service.accountsCollectionName)
		if err != nil {
			return err
		}
		service.accountsCollection = collection
	}
	return service.accountsCollection.FindAll(nil, result)
}

// UpsertAccount creates or updates an account.
func (service *CosmosAccountsService) UpsertAccount(ctx context.Context, account *Account) error {
	if service.accountsCollection == nil {
		collection, err := service.collectionFactory.CreateCollection(ctx, service.accountsCollectionName)
		if err != nil {
			return err
		}
		service.accountsCollection = collection
	}
	_, err := service.accountsCollection.Upsert(bson.M{"shardkey": account.Address}, account)
	return err
}
