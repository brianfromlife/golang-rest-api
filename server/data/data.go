package data

import (
	"github.com/brianfromlife/golang-ecs/server/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDataProvider interface {
	IUserProvider
	IAccountProvider
}

type DataProvider struct {
	db *mongo.Client
}

func New(cfg *config.Settings, db *mongo.Client) IDataProvider {
	return DataProvider{}
}
