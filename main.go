package main

import (
	"github.com/brianfromlife/golang-ecs/server"

	"github.com/brianfromlife/golang-ecs/server/config"
	"github.com/brianfromlife/golang-ecs/server/db"
)

func main() {
	cfg := config.New()
	db := db.NewConnection(cfg)
	defer db.Client.Disconnect(db.Ctx)
	svr := server.New(cfg, db.Client)
	svr.Start()
}
