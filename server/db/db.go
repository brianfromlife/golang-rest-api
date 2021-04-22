package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianfromlife/golang-ecs/server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	Ctx    context.Context
}

func NewConnection(cfg *config.Settings) Connection {

	uri := "mongodb://localhost:27017/golang_ecs"

	credentials := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPassword,
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

	return Connection{
		Client: client,
		Ctx:    ctx,
	}
}
