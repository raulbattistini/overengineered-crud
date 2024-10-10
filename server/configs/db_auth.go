package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() *mongo.Client {
	connectionString := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Println(err) // proper logging should be added
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Println(err)
		}
	}()

	return client
}
