package api

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server  *echo.Echo
	userSvc services.IUserService
	cfg     *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, userProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		cfg:     cfg,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
	a.server.POST("/v1/public/account/login", a.Login)

	protected := a.server.Group("v1/api")

	middleware := Middleware{config: a.cfg}

	protected.Use(middleware.Auth)
	protected.GET("/secret", func(c echo.Context) error {
		userId := c.Get("user").(string)
		return c.String(200, userId)
	})
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
