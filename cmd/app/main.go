package main

import (
	"log"

	"gahmen-api/config"
	"gahmen-api/db"
	"gahmen-api/cmd/server"
)

func main() {
	config := config.NewConfig()
	config.Parse()
	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	store, err := storage.NewPostgresStore(config)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":3080", store, config)
	server.Run()
}
