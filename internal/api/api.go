package api

import (
	"github.com/brianfromlife/golang-ecs/internal/api/handlers"
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/logger"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server *echo.Echo
}

func Start(cfg *config.Settings, client *mongo.Client) {
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())

	// logger
	logger := logger.NewLogger(cfg)

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, logger, userProvider)

	// handlers
	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler(userSvc)

	server.GET("/v1/public/healthy", healthHandler.HealthCheck)
	server.POST("/v1/public/account/register", userHandler.Register)

	server.Start(":5000")
}
