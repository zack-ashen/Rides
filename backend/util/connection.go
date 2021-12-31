package util

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "rides"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		
		// Connect to MongoDB
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
