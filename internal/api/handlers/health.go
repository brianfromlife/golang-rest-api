package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (hh HealthHandler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
