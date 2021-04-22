package server

import (
	"github.com/brianfromlife/golang-ecs/server/config"
	"github.com/brianfromlife/golang-ecs/server/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server *echo.Echo
	data   data.IDataProvider
}

func New(cfg *config.Settings, db *mongo.Client) App {
	server := echo.New()
	data := data.New(cfg, db)
	return App{
		server: server,
		data:   data,
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
