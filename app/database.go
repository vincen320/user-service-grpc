package app

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	clientOptions := options.Client().
		SetMaxPoolSize(100).
		SetMaxConnecting(20).
		SetConnectTimeout(60 * time.Minute).
		SetMaxConnIdleTime(10 * time.Minute).
		ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client.Database("v_user_grpc"), nil
}
