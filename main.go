package main

import (
	"github.com/brianfromlife/golang-ecs/api"

	"github.com/brianfromlife/golang-ecs/api/config"
	"github.com/brianfromlife/golang-ecs/api/data"
)

func main() {
	cfg := config.New()
	db := data.NewConnection(cfg)
	defer db.Client.Disconnect(db.Ctx)
	app := api.New(cfg, db.Client)
	app.Start()
}
