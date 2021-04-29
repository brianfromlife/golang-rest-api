package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	server *echo.Echo
}

func New() App {
	server := echo.New()

	return App{
		server: server,
	}
}

func (a App) configureGlobalMiddleware() {
	a.server.Use(middleware.Recover())

}

func (a App) setupRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
}

func (a App) Start() {
	a.configureGlobalMiddleware()
	a.setupRoutes()
	a.server.Logger.Fatal(a.server.Start(":5000"))
}
