package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {
	clientOptions := options.Client().
		SetMaxPoolSize(100).
		SetMaxConnecting(20).
		SetConnectTimeout(60 * time.Minute).
		SetMaxConnIdleTime(10 * time.Minute).
		ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("cannot connect mongodb", err)
	}
	return client.Database("v_user_grpc")
}
