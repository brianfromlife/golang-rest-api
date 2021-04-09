package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

func New() mongoClient {

	uri := "mongodb://localhost:27017/golang_ecs"

	credentials := options.Credential{
		Username: "root",
		Password: "password",
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database.")

	return mongoClient{
		client: client,
		ctx:    ctx,
	}
}

func (mc mongoClient) DisconnectClient() {
	mc.client.Disconnect(mc.ctx)
}
