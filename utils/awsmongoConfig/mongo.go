package awsmongoconfig

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDBInstance(connectionString, dbName string) (*MongoDB, error) {
	fmt.Print(connectionString, dbName)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	err = client.Ping(context.Background(), nil) //checking the connection is working or not.
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}
	db := client.Database(dbName)
	instance := &MongoDB{
		Client:   client,
		Database: db,
	}
	// fmt.Print(instance)
	return instance, nil
}

// Close closes the MongoDB client connection.
func (m *MongoDB) Close() error {
	return m.Client.Disconnect(context.Background())
}
