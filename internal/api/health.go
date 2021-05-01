package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a App) HealthCheck(c echo.Context) error {
	a.logger.Info("testing")
	return c.String(http.StatusOK, "healthy")
}
