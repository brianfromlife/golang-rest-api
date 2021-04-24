package data

import (
	"github.com/brianfromlife/golang-ecs/api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDataProvider interface {
	IUserProvider
}

type DataProvider struct {
	db *mongo.Database
}

func New(cfg *config.Settings, mongo *mongo.Client) IDataProvider {
	return DataProvider{db: mongo.Database(cfg.DbName)}
}
