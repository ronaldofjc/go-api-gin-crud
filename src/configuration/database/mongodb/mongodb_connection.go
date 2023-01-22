package mongodb

import (
	"api-crud/src/configuration/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitConnection() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("MongoDB connected with success")
}
