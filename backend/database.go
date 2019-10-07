//This file contains functions for database
package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(cs string) *mongo.Client {
	log.WithField("connection-string", cs).Info("Connecting to mongo-db")
	clientOptions := options.Client().ApplyURI(cs)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.WithField("err", err).Fatal("Can't Connect to mongodb server")
	}

	return client
}

func DisconnectClient(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err == nil {
		log.Info("MongoDB disconnected")
	}
}
