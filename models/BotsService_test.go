package models

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/globalsign/mgo/bson"
)

func TestMockBotsService_CreateBotState(t *testing.T) {
	mockCollection := &MockCollection{}
	stateCollectionName := "mock-bots"
	type fields struct {
		CollectionFactory   CollectionFactory
		stateCollectionName string
	}
	type wants struct {
		stateCollection     Collection
		stateCollectionName string
		state               BotState
	}
	tests := []struct {
		name    string
		fields  fields
		wants   wants
		wantErr bool
	}{
		{
			name: "Should create bot state.",
			fields: fields{
				CollectionFactory: &MockCollectionFactory{
					Collection: mockCollection,
				},
				stateCollectionName: stateCollectionName,
			},
			wants: wants{
				stateCollection:     mockCollection,
				stateCollectionName: stateCollectionName,
				state:               &MockBotState{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &MockBotsService{
				CollectionFactory:   tt.fields.CollectionFactory,
				stateCollectionName: tt.fields.stateCollectionName,
			}
			ctx := context.Background()
			err := service.CreateBotState(ctx, tt.wants.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBotState(ctx, tt.wants.state) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(mockCollection, tt.wants.stateCollection) {
				t.Errorf("mockCollection = %v, want %v", mockCollection, tt.wants.stateCollection)
			}
			name := mockCollection.GetName()
			if name != tt.wants.stateCollectionName {
				t.Errorf("mockCollection.GetName()= %v, want %v", name, tt.wants.stateCollectionName)
			}
			got := &MockBotState{}
			err = mockCollection.FindOne(nil, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockCollection.FindOne(nil) error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wants.state) {
				t.Errorf("got, _ = mockCollection.FindOne(nil) = %v, want %v", got, tt.wants.state)
			}
		})
	}
}

func TestDocumentDbBotsService_CreateBotState(t *testing.T) {
	cosmosClient := &CosmosClient{
		logger: log.New(os.Stderr, "CosmosClient | ", 0),
		configuration: CosmosConfiguration{
			SubscriptionID:    "6951e94d-0947-4c5e-b865-f864609da246",
			CloudName:         "AzurePublicCloud",
			ResourceGroupName: "crmbts-devmachines-bao-blockchain-460259",
			AccountName:       "bao-blockchain",
		},
	}
	cosmosClient.Connect()
	ctx := context.Background()
	session, err := cosmosClient.GetSession(ctx)
	if err != nil {
		t.Errorf("cannot get mongoDB session: %v", err)
	}
	defer session.Close()
	documentDbCollectionFactory := &CosmosCollectionFactory{
		logger:       log.New(os.Stderr, "DocumentDbCollectionFactory | ", 0),
		cosmosClient: cosmosClient,
		session:      session,
	}
	type fields struct {
		logger              *log.Logger
		collectionFactory   CollectionFactory
		stateCollectionName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    BotState
		wantErr bool
	}{
		{
			name: "Should create bot state.",
			fields: fields{
				logger:              log.New(os.Stderr, "DocumentDbBotsService | ", 0),
				collectionFactory:   documentDbCollectionFactory,
				stateCollectionName: "mock-bots",
			},
			want: &MockBotState{
				ShardKey: bson.NewObjectId().Hex(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &CosmosBotsService{
				logger:              tt.fields.logger,
				collectionFactory:   tt.fields.collectionFactory,
				stateCollectionName: tt.fields.stateCollectionName,
			}
			ctx := context.Background()
			err := service.CreateBotState(ctx, tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBotState(ctx, state) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := &MockBotState{}
			err = service.stateCollection.FindOne(nil, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.stateCollection.FindOne(nil) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.GetShardKey() != tt.want.GetShardKey() {
				t.Errorf(`got.GetShardKey() = %v, want %v`, got.GetShardKey(), tt.want.GetShardKey())
			}
			err = cosmosClient.DeleteSQLContainer(ctx, tt.fields.stateCollectionName)
			if err != nil {
				t.Errorf("Failed to delete container: %v", err)
			}
		})
	}
}
