package db

import (
	"context"
	"fmt"
	"server/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func GetConnection(config configs.MongoDBConfig) (*mongo.Client, error) {
	connectionString := fmt.Sprintf("%s%s:%s", config.Protocol, config.Host, config.Port)

	mongoOptions := options.Client().ApplyURI(connectionString)
	mongoConnection, err := mongo.Connect(context.Background(), mongoOptions)

	if err != nil {
		return nil, err
	}
	defer func() {
		if err = mongoConnection.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	MongoDB = mongoConnection

	return mongoConnection, nil
}
