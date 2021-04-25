package main

import (
	"github.com/brianfromlife/golang-ecs/api"
	"github.com/brianfromlife/golang-ecs/api/db"

	"github.com/brianfromlife/golang-ecs/api/config"
)

func main() {
	cfg := config.New()
	db := db.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
