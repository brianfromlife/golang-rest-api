package main

import (
	"github.com/brianfromlife/golang-ecs/server"

	"github.com/brianfromlife/golang-ecs/server/db"
)

func main() {
	db := db.New()
	defer db.DisconnectClient()

	svr := server.New()
	svr.Start()
}
