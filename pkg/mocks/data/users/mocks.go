package usermocks

import (
	"github.com/brianfromlife/golang-ecs/pkg/domain"
)

type UserProviderMock struct {
	CreateAcountMock   func(user *domain.User) error
	UsernameExistsMock func(username string) (bool, error)
	FindByUsernameMock func(username string) (*domain.User, error)
}

func (u UserProviderMock) CreateAcount(user *domain.User) error {
	return u.CreateAcountMock(user)
}

func (u UserProviderMock) UsernameExists(username string) (bool, error) {
	return u.UsernameExistsMock(username)
}

func (u UserProviderMock) FindByUsername(username string) (*domain.User, error) {
	return u.FindByUsernameMock(username)
}
