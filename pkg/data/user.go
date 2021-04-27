package data

import (
	"context"

	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserProvider interface {
	CreateAcount(user *models.User) error
	UsernameExists(username string) (bool, error)
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

func (u UserProvider) UsernameExists(username string) (bool, error) {
	var userFound models.User
	filter := bson.D{primitive.E{Key: "text", Value: username}}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding by username")
	}

	return true, nil
}
