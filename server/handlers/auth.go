package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (h AuthHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
