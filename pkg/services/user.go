package user

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
)

type IUserService interface{}

type UserService struct {
	userProvider data.IUserProvider
}

func NewUserService(cfg *config.Settings, userProvider data.IUserProvider) IUserService {
	return UserService{
		userProvider: userProvider,
	}
}
