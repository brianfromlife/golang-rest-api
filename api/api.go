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
	Data   data.IDataProvider
	cfg    *config.Settings
	logger logger.ILogger
}

func New(cfg *config.Settings, db *mongo.Client) App {
	server := echo.New()
	// data := data.New(cfg, db)
	return App{
		server: server,
		// Data:   data,
		cfg: cfg,
	}
}

func (a App) ConfigureRoutes() {
	public := a.server.Group("/v1/public")
	public.GET("/healthy", a.HealthCheck)
	public.POST("/account/register", a.RegisterAccount)

	private := a.server.Group("/v1")
	private.Use(ServerHeader)
	private.GET("/test", func(c echo.Context) error {
		return c.String(200, "hhwe")
	})
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Logger.Fatal(a.server.Start(":5000"))
}
