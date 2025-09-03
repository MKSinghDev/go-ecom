package main

import (
	"log"

	"github.com/MKSinghDev/go-ecom/src/api"
	"github.com/MKSinghDev/go-ecom/src/config"
	"github.com/MKSinghDev/go-ecom/src/db"
)

func main() {
	dbpool := db.NewPsqlStorage(config.Envs.DBConnString)
	db.InitStorage(dbpool)

	server := api.NewAPIServer(config.Envs.Port, dbpool)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
