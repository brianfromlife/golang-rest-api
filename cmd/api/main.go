package main

import (
	"github.com/brianfromlife/golang-ecs/internal/api"
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
)

func main() {
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
