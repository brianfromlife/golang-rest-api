package api

import (
	"net/http"

	"github.com/brianfromlife/golang-ecs/pkg/models"
	"github.com/labstack/echo/v4"
)

func (a App) Register(c echo.Context) error {
	newUser, err := models.ValidateRegisterRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	err = a.userSvc.CreateAccount(newUser)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusOK, "register")
}
