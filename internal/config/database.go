package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectMongodb(mongoHost string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoOpts := options.Client().ApplyURI(mongoHost).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return mongoClient
}

func CloseDatabaseConnection(mongoClient *mongo.Client) {
	if mongoClient != nil {
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Disconnected to MongoDB")
	}
}
