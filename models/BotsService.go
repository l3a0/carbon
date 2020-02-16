package models

import (
	"context"
	"log"
)

// BotsService is responsible for CRUD on bot data.
type BotsService interface {
	// CreateCollection(ctx context.Context) (Collection, error)
	CreateBotState(ctx context.Context, state BotState) error
}

// BotState is responsible for accessing bot state.
type BotState interface {
	GetShardKey() string
}

// MockBotState stores bot state.
type MockBotState struct {
	ShardKey string
}

// GetShardKey returns the shard key.
func (state *MockBotState) GetShardKey() string {
	return state.ShardKey
}

// MockBotsService works against an in-memory data store that is not durable.
type MockBotsService struct {
	CollectionFactory   CollectionFactory
	stateCollection     Collection
	stateCollectionName string
}

// CreateCollection initializes the in-memory data store.
func (service *MockBotsService) CreateCollection(ctx context.Context) (Collection, error) {
	collection, err := service.CollectionFactory.CreateCollection(ctx, service.stateCollectionName)
	service.stateCollection = collection
	return collection, err
}

// CreateBotState creates the bot state in the collection.
func (service *MockBotsService) CreateBotState(ctx context.Context, state BotState) error {
	if service.stateCollection == nil {
		_, err := service.CreateCollection(ctx)
		if err != nil {
			return err
		}
	}
	return service.stateCollection.Create(state)
}

// CosmosBotsService works against Cosmos DB SQL Core.
type CosmosBotsService struct {
	logger              *log.Logger
	collectionFactory   CollectionFactory
	stateCollection     Collection
	stateCollectionName string
}

// CreateBotState creates the bot state in the collection.
func (service *CosmosBotsService) CreateBotState(ctx context.Context, state BotState) error {
	if service.stateCollection == nil {
		service.logger.Printf("State Collection not found: %v.\n", service.stateCollectionName)
		service.logger.Printf("Creating State Collection: %v.\n", service.stateCollectionName)
		collection, err := service.collectionFactory.CreateCollection(ctx, service.stateCollectionName)
		if err != nil {
			service.logger.Printf("Failed to create State Collection: %v.\n", service.stateCollectionName)
			return err
		}
		service.stateCollection = collection
		service.logger.Printf("Created State Collection: %v.\n", service.stateCollection.GetName())
	}
	service.logger.Printf("Creating Bot State: %v.\n", state)
	err := service.stateCollection.Create(state)
	service.logger.Printf("Created Bot State: %v.\n", state)
	return err
}
