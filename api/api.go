package api

import (
	"github.com/brianfromlife/golang-ecs/api/config"
	"github.com/brianfromlife/golang-ecs/api/data"
	"github.com/brianfromlife/golang-ecs/api/logger"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server *echo.Echo
	data   data.IDataProvider
	cfg    *config.Settings
	logger logger.ILogger
}

func New(cfg *config.Settings, db *mongo.Client) App {
	server := echo.New()
	data := data.New(cfg, db)
	return App{
		server: server,
		data:   data,
		cfg:    cfg,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.RegisterAccount)
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Logger.Fatal(a.server.Start(":5000"))
}
