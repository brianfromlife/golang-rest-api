package server

import "github.com/labstack/echo/v4"

type app struct {
	server *echo.Echo
}

func New() app {
	server := echo.New()
	return app{
		server: server,
	}
}

func (a app) Start() {
	a.RouterSetup()
	a.server.Logger.Fatal(a.server.Start(":5000"))
}
