package services

import (
	"testing"

	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/domain"
	usermocks "github.com/brianfromlife/golang-ecs/pkg/mocks/data/users"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount_UserExists(t *testing.T) {
	cfg := &config.Settings{}
	userProviderMock := &usermocks.UserProviderMock{}
	userProviderMock.UsernameExistsMock = func(username string) (bool, error) {
		return true, nil
	}

	userSvc := NewUserService(cfg, userProviderMock)
	newUser := &domain.User{Username: "test", Password: "test"}

	response := userSvc.CreateAccount(newUser)
	assert.Equal(t, "USERNAME_TAKEN", response.Name)

}
