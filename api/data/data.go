package data

import (
	"context"

	"github.com/brianfromlife/golang-ecs/api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDataProvider interface {
	IUserProvider
}

type DataProvider struct {
	userCollection *mongo.Collection
	todoCollection *mongo.Collection
	ctx            context.Context
}

func New(cfg *config.Settings, mongo *mongo.Client) IDataProvider {
	return DataProvider{userCollection: mongo.Database(cfg.DbName).Collection("users"), ctx: context.TODO()}
}
