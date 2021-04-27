package user

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	CreateAccount(user *models.User) error
}

type UserService struct {
	userProvider data.IUserProvider
}

func NewUserService(cfg *config.Settings, userProvider data.IUserProvider) IUserService {
	return UserService{
		userProvider: userProvider,
	}
}

func (u UserService) CreateAccount(user *models.User) error {
	user.ID = primitive.NewObjectID()
	userExists, err := u.userProvider.UsernameExists(user.Username)

	if err != nil {
		return err
	}

	if userExists {
		return err
	}

	err = u.userProvider.CreateAcount(user)

	if err != nil {
		return err
	}

	return nil
}

func hashPassword() {

}
