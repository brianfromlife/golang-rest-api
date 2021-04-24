package api

import (
	"net/http"

	"github.com/brianfromlife/golang-ecs/api/requests"
	"github.com/labstack/echo/v4"
)

func (a App) RegisterAccount(c echo.Context) error {

	_, err := requests.ValidateRegisterAccount(c)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusOK, "register account")
}
