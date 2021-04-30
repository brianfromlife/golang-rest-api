package handlers

import (
	"net/http"

	"github.com/brianfromlife/golang-ecs/pkg/models"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userSvc services.IUserService
}

func NewUserHandler(userSvc services.IUserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (h UserHandler) Register(c echo.Context) error {
	newUser, err := models.ValidateRegisterRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	err = h.userSvc.CreateAccount(newUser)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusOK, "register")
}
