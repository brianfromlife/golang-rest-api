package usermocks

import (
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/domain"
)

var mockData = []domain.User{
	{
		Username: "user1",
		Password: "password1",
	},
	{
		Username: "user2",
		Password: "password2",
	},
}

type MockDataStock struct{}

func NewUserDataStoreMock() data.IUserProvider {
	return &MockDataStock{}
}

func (u MockDataStock) CreateAcount(user *domain.User) error {
	mockData = append(mockData, *user)
	return nil
}

func (u MockDataStock) UsernameExists(username string) (bool, error) {
	for user := range mockData {
		if mockData[user].Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (u MockDataStock) FindByUsername(username string) (*domain.User, error) {
	for user := range mockData {
		if mockData[user].Username == username {
			return &mockData[user], nil
		}
	}
	return nil, nil
}
