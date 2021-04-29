package main

import (
	"github.com/brianfromlife/golang-ecs/internal/api"
	"github.com/brianfromlife/golang-ecs/pkg/data"
)

func main() {
	db := data.NewMongoConnection()
	defer db.Disconnect()
	application := api.New()
	application.Start()
}
