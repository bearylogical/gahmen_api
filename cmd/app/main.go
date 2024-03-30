package main

import (
	"log"

	"gahmen-api/cmd/server"
	"gahmen-api/config"
	storage "gahmen-api/db"
)

func main() {
	log.Println("Starting Gahmen API server...")
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
