package models

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/globalsign/mgo"
)

// OAuthGrantType specifies which grant type to use.
type OAuthGrantType int

const (
	// OAuthGrantTypeServicePrincipal for client credentials flow
	OAuthGrantTypeServicePrincipal OAuthGrantType = iota
	// OAuthGrantTypeDeviceFlow for device flow
	OAuthGrantTypeDeviceFlow
)

// CosmosService is responsible for CRUD on Cosmos DB.
type CosmosService interface {
	Connect()
	GetDatabaseName() string
	GetDatabaseClient() documentdb.DatabaseAccountsClient
	GetSession(ctx context.Context) (session *mgo.Session, err error)
	GetSQLContainer(ctx context.Context, containerName string) (result documentdb.SQLContainer, err error)
	CreateUpdateSQLContainer(ctx context.Context, containerName string, parameters documentdb.SQLContainerCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateSQLContainerFuture, err error)
	DeleteSQLContainer(ctx context.Context, containerName string) (err error)
}

// CosmosConfiguration contains Cosmos DB configuration.
type CosmosConfiguration struct {
	SubscriptionID    string
	CloudName         string
	ResourceGroupName string
	AccountName       string
}

// CosmosClient stores Cosmos DB client.
type CosmosClient struct {
	logger        *log.Logger
	configuration CosmosConfiguration
	client        documentdb.DatabaseAccountsClient
}

// NewCosmosService creates a new CosmosService
func NewCosmosService(logger *log.Logger, configuration CosmosConfiguration) CosmosService {
	return &CosmosClient{
		logger:        logger,
		configuration: configuration,
	}
}

// Connect initializes a new client.
func (c *CosmosClient) Connect() {
	c.client = documentdb.NewDatabaseAccountsClient(c.configuration.SubscriptionID)
	c.client.PollingDelay = 5 * time.Second
	// create an authorizer from Azure Managed Service Idenity
	auth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	c.client.Authorizer = auth
	c.client.AddToUserAgent("carbon-tests")
}

// GetDatabaseName returns the database name.
func (c *CosmosClient) GetDatabaseName() string {
	return c.configuration.AccountName
}

// GetDatabaseClient returns the database client.
func (c *CosmosClient) GetDatabaseClient() documentdb.DatabaseAccountsClient {
	return c.client
}

// GetSession returns a new session.
func (c *CosmosClient) GetSession(ctx context.Context) (session *mgo.Session, err error) {
	keys, err := c.client.ListKeys(ctx, c.configuration.ResourceGroupName, c.configuration.AccountName)
	if err != nil {
		c.logger.Printf("cannot list keys: %v\n", err)
		return nil, err
	}
	host := fmt.Sprintf("%s.documents.azure.com", c.configuration.AccountName)
	session, err = mongodb.NewMongoDBClientWithCredentials(c.configuration.AccountName, *keys.PrimaryMasterKey, host)
	if err != nil {
		c.logger.Printf("cannot get mongoDB session: %v", err)
		return nil, err
	}
	return session, nil
}

// GetSQLContainer gets the SQL container under an existing Azure Cosmos DB database account.
// Parameters:
// containerName - cosmos DB container name.
func (c *CosmosClient) GetSQLContainer(ctx context.Context, containerName string) (result documentdb.SQLContainer, err error) {
	return c.client.GetSQLContainer(ctx, c.configuration.ResourceGroupName, c.configuration.AccountName, c.configuration.AccountName, containerName)
}

// CreateUpdateSQLContainer create or update an Azure Cosmos DB SQL container
// Parameters:
// containerName - cosmos DB container name.
// createUpdateSQLContainerParameters - the parameters to provide for the current SQL container.
func (c *CosmosClient) CreateUpdateSQLContainer(ctx context.Context, containerName string, parameters documentdb.SQLContainerCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateSQLContainerFuture, err error) {
	return c.client.CreateUpdateSQLContainer(ctx, c.configuration.ResourceGroupName, c.configuration.AccountName, c.configuration.AccountName, containerName, parameters)
}

// DeleteSQLContainer deletes an existing Azure Cosmos DB SQL container.
// Parameters:
// containerName - cosmos DB container name.
func (c *CosmosClient) DeleteSQLContainer(ctx context.Context, containerName string) (err error) {
	future, err := c.client.DeleteSQLContainer(ctx, c.configuration.ResourceGroupName, c.configuration.AccountName, c.configuration.AccountName, containerName)
	if err != nil {
		return err
	}
	err = future.WaitForCompletionRef(ctx, c.GetDatabaseClient().Client)
	if err != nil {
		return err
	}
	response, err := future.Result(c.GetDatabaseClient())
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	c.logger.Printf("Deleted container: %v", string(buf))
	return nil
}

// Collection is responsible for CRUD on data.
type Collection interface {
	FindOne(query interface{}, result interface{}) error
	FindAll(query interface{}, result interface{}) error
	Create(doc interface{}) error
	Update(selector interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	GetName() string
}

// CollectionFactory is responsible for creating collections.
type CollectionFactory interface {
	CreateCollection(ctx context.Context, name string) (Collection, error)
}

// MockCollectionFactory is responsible for creating in-memory collections.
type MockCollectionFactory struct {
	Collection *MockCollection
}

// CreateCollection initializes the in-memory data store.
func (factory *MockCollectionFactory) CreateCollection(ctx context.Context, name string) (Collection, error) {
	factory.Collection.name = name
	return factory.Collection, nil
}

// MockCollection is an in-memory data store that is not durable.
type MockCollection struct {
	name string
	data []interface{}
}

// FindOne returns the first record for testing.
func (collection *MockCollection) FindOne(query interface{}, result interface{}) error {
	result = collection.data[0]
	return nil
}

// FindAll returns the collection.
func (collection *MockCollection) FindAll(query interface{}, result interface{}) error {
	result = collection.data
	return nil
}

// Create the record.
func (collection *MockCollection) Create(doc interface{}) error {
	collection.data = append(collection.data, doc)
	return nil
}

// Update the record.
func (collection *MockCollection) Update(selector interface{}, update interface{}) error {
	collection.data[0] = update
	return nil
}

// Upsert updates or creates the record.
func (collection *MockCollection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	collection.data[0] = update
	return nil, nil
}

// GetName returns the collection name.
func (collection *MockCollection) GetName() string {
	return collection.name
}

// CosmosCollection is responsible for adapting to Collection interface.
type CosmosCollection struct {
	collection *mgo.Collection
}

// FindOne returns the first record matching the query.
func (adapter *CosmosCollection) FindOne(query interface{}, result interface{}) error {
	return adapter.collection.Find(query).One(result)
}

// FindAll returns all records matching the query.
func (adapter *CosmosCollection) FindAll(query interface{}, result interface{}) error {
	return adapter.collection.Find(query).All(result)
}

// Create the records
func (adapter *CosmosCollection) Create(doc interface{}) error {
	return adapter.collection.Insert(doc)
}

// Update the record.
func (adapter *CosmosCollection) Update(selector interface{}, update interface{}) error {
	return adapter.collection.Update(selector, update)
}

// Upsert updates or creates the record.
func (adapter *CosmosCollection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return adapter.collection.Upsert(selector, update)
}

// GetName returns the collection name.
func (adapter *CosmosCollection) GetName() string {
	return adapter.collection.Name
}

// CosmosCollectionFactory is responsible for creating Document DB collections.
type CosmosCollectionFactory struct {
	logger       *log.Logger
	cosmosClient CosmosService
	session      *mgo.Session
}

// NewCosmosCollectionFactory creates a new CollectionFactory.
func NewCosmosCollectionFactory(logger *log.Logger, cosmosClient CosmosService, session *mgo.Session) CollectionFactory {
	return &CosmosCollectionFactory{
		logger:       logger,
		cosmosClient: cosmosClient,
		session:      session,
	}
}

// CreateCollection initializes the in-memory data store.
func (factory *CosmosCollectionFactory) CreateCollection(ctx context.Context, collectionName string) (Collection, error) {
	adapter := &CosmosCollection{}
	_, err := factory.cosmosClient.GetSQLContainer(ctx, collectionName)
	if err != nil {
		factory.logger.Printf("Collection not found: %v.\n", collectionName)
	} else {
		adapter.collection = factory.session.DB(factory.cosmosClient.GetDatabaseName()).C(collectionName)
		factory.logger.Printf("Collection found: %v.\n", adapter.GetName())
		return adapter, nil
	}
	containerParameters := documentdb.SQLContainerCreateUpdateParameters{
		SQLContainerCreateUpdateProperties: &documentdb.SQLContainerCreateUpdateProperties{
			Resource: &documentdb.SQLContainerResource{
				ID: &collectionName,
				PartitionKey: &documentdb.ContainerPartitionKey{
					Paths: &[]string{"/shardkey"},
					Kind:  documentdb.PartitionKindHash,
				},
			},
			Options: map[string]*string{},
		},
	}
	factory.logger.Printf("Creating collection: %v.\n", *containerParameters.Resource.ID)
	future, err := factory.cosmosClient.CreateUpdateSQLContainer(ctx, collectionName, containerParameters)
	if err != nil {
		return nil, err
	}
	err = future.WaitForCompletionRef(ctx, factory.cosmosClient.GetDatabaseClient().Client)
	if err != nil {
		return nil, err
	}
	_, err = future.Result(factory.cosmosClient.GetDatabaseClient())
	if err != nil {
		return nil, err
	}
	adapter.collection = factory.session.DB(factory.cosmosClient.GetDatabaseName()).C(collectionName)
	factory.logger.Printf("Created collection: %v.\n", adapter.GetName())
	return adapter, nil
}
