package api

import (
	"net/http"

	"github.com/brianfromlife/golang-ecs/api/models"
	"github.com/brianfromlife/golang-ecs/api/requests"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a App) RegisterAccount(c echo.Context) error {

	user, err := requests.ValidateRegisterAccount(c)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	newUser := &models.User{
		ID:       primitive.NewObjectID(),
		Username: user.Username,
		Password: user.Password,
	}

	error := a.Data.CreateAcount(newUser)

	if error != nil {
		a.logger.Error("error creating user", error)
		return c.JSON(http.StatusBadRequest, "There was an error creating your account. Please try again.")
	}

	return c.String(http.StatusCreated, "")
}
