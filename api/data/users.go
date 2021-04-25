package data

import (
	"github.com/brianfromlife/golang-ecs/api/models"
)

type IUserProvider interface {
	CreateAcount(user *models.User) error
	FindUserById() string
}

func (d DataProvider) CreateAcount(user *models.User) error {
	_, err := d.userCollection.InsertOne(d.ctx, user)
	return err
}

func (d DataProvider) FindUserById() string {
	return "string"
}
