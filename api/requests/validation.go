package requests

import (
	"github.com/brianfromlife/golang-ecs/server/errors"
	"github.com/brianfromlife/golang-ecs/server/models"
	"github.com/labstack/echo/v4"
)

func ValidateRegisterAccount(c echo.Context) (*models.NewUser, *errors.ApiError) {
	newUserRequest := new(RegisterAccountRequest)
	if err := c.Bind(newUserRequest); err != nil {
		return nil, errors.BindError()
	}

	var validationErrors []string

	if len(newUserRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(newUserRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(validationErrors) > 0 {
		return nil, errors.ValidationError(validationErrors)
	}

	return &models.NewUser{
		Username: newUserRequest.Username,
		Password: newUserRequest.Password,
	}, nil
}
