package api

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	server *echo.Echo
	logger logger.ILogger
	cfg    *config.Settings
}

func New(cfg *config.Settings) App {
	server := echo.New()

	return App{
		server: server,
		cfg:    cfg,
	}
}

func (a App) configureLogger() {
	if a.cfg.Env == "development" {
		logger := logger.NewLocal()
		a.logger = logger
	} else {
		logger := logger.NewLogger("secret")
		a.logger = logger
	}
}

func (a App) configureGlobalMiddleware() {
	a.server.Use(middleware.Recover())

}

func (a App) setupRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
}

func (a App) Start() {
	a.configureLogger()
	a.configureGlobalMiddleware()
	a.setupRoutes()
	a.server.Logger.Fatal(a.server.Start(":5000"))
}
