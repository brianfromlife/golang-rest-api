package user

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateAccount(user *models.User) error
	Login() (string, error)
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
	userExists, err := u.userProvider.UsernameExists(user.Username)

	if err != nil {
		return err
	}

	if userExists {
		return err
	}

	user.ID = primitive.NewObjectID()
	hash, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hash

	err = u.userProvider.CreateAcount(user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserService) Login() (string, error) {
	return "", nil
}

func hashPassword(password string) (string, error) {

	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 12)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func comparePasswordWithHash(password string, hash string) error {

	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	return err
}
