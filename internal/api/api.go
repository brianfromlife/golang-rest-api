package api

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/logger"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server  *echo.Echo
	logger  logger.ILogger
	userSvc services.IUserService
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())

	// logger
	logger := logger.NewLogger(cfg)

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, logger, userProvider)

	return &App{
		server:  server,
		logger:  logger,
		userSvc: userSvc,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
