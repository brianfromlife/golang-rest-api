package data

import (
	"context"

	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserProvider interface {
	CreateAcount(user *models.User) error
	FindByUsername() string
}

type UserProvider struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserProvider(cfg *config.Settings, mongo *mongo.Client) IUserProvider {
	userCollection := mongo.Database(cfg.DbName).Collection("users")
	return UserProvider{
		userCollection: userCollection,
		ctx:            context.TODO(),
	}
}

func (u UserProvider) CreateAcount(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u UserProvider) FindByUsername() string {
	return "string"
}
