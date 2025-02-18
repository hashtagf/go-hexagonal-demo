package mongodb

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (*mongo.Client, error) {
	ctx := context.Background()
	var mongoURI string

	if os.Getenv("MODE") == "dev" {
		mongoURI = "mongodb://localhost:27017"
	} else if os.Getenv("MODE") == "prod" {
		mongoURI = os.Getenv("MONGO_URI")
	} else {
		return nil, errors.New("invalid mode")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
