package database

import (
	"context"
	"errors"
	"fmt"
	"sync"

	// "github.com/sadaghiani/concurrent-file-processing/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type IMongoDataStore interface {
	GetDatabase() *mongo.Database
	GetClient() (*mongo.Client, error)
	Coll(name string, opts ...*options.CollectionOptions) *mongo.Collection
	Disconnect() error
}

func (md *MongoDataStore) GetDatabase() *mongo.Database {
	return md.database
}

func (md *MongoDataStore) GetClient() (*mongo.Client, error) {
	if md.client != nil {
		return md.client, nil
	}
	return nil, errors.New("client is missing (nil) in Mongo Data Store")
}

func (md *MongoDataStore) Coll(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return md.database.Collection(name, opts...)
}

func (md *MongoDataStore) Disconnect() error {
	err := md.client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

type MongoDataStore struct {
	database *mongo.Database
	client   *mongo.Client
	logger   *zap.Logger
	config   MongoDataStoreConfig
}

type MongoDataStoreConfig struct {
	DatabaseURI,
	DatabasePort,
	DatabaseName,
	DatabaseCollection,
	DatabaseUser,
	DatabasePassword string
}

func NewMongoDataStore(logger *zap.Logger, mongoDataStoreConfig MongoDataStoreConfig) *MongoDataStore {

	md := MongoDataStore{
		logger: logger,
		config: mongoDataStoreConfig,
	}
	md.connect()
	if md.database != nil && md.client != nil {
		return &md
	}
	md.logger.Fatal("Failed to connect to database")
	return nil
}

func (md *MongoDataStore) connect() {
	var connectOnce sync.Once
	connectOnce.Do(func() {
		md.database, md.client = md.connectToMongo()
	})
}

func (md *MongoDataStore) connectToMongo() (*mongo.Database, *mongo.Client) {
	var err error
	client, err := mongo.NewClient(md.newClientOption())
	if err != nil {
		md.logger.Fatal("cannot create newClient", zap.Error(err))
	}
	client.Connect(context.TODO())
	if err != nil {
		md.logger.Fatal("cannot client connect", zap.Error(err))
	}
	var database = client.Database(md.config.DatabaseName)
	md.logger.Info("conncected to database")
	return database, client
}

func (md *MongoDataStore) newClientOption() *options.ClientOptions {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      md.config.DatabaseUser,
		Password:      md.config.DatabasePassword,
	}
	uri := fmt.Sprintf("mongodb://%s:%s/", md.config.DatabaseURI, md.config.DatabasePort)
	return options.Client().ApplyURI(uri).SetAuth(credential)
}
