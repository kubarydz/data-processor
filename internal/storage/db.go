package storage

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dsn_default = "mongodb://localhost:27017"

func ProvideCollection() (*mongo.Collection, error) {
	dsn := os.Getenv("MONGO_DSN")
	if dsn == "" {
		dsn = dsn_default
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	return client.Database("users").Collection("users"), err
}
