package main

import (
	"github.com/brianfromlife/golang-ecs/api"

	"github.com/brianfromlife/golang-ecs/api/config"
	"github.com/brianfromlife/golang-ecs/api/data"
)

func main() {
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
