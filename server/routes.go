package server

import "github.com/brianfromlife/golang-ecs/server/handlers"

func (a app) RouterSetup() {
	healthHandler := handlers.NewHealthHandler()
	a.server.GET("/public/healthy", healthHandler.Healthy)
}
