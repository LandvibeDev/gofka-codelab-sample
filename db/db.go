package db

import (
	"context"
	"github.com/LandvibeDev/gofka-codelab-sample/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(config config.DatabaseConfiguration) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Hosts)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
