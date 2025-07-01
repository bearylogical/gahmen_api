package main

import (
	"log"

	"gahmen-api/cmd/server"
	"gahmen-api/config"
	storage "gahmen-api/db"
)

// @title           Gahmen Budget API
// @version         1.0
// @description     Gahmen Budget API provides access to Singapore's government budget data.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Syamil Maulod
// @contact.url    https://bearylogical.net
// @contact.email  syamil@bearylogical.net

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1
// @schemes   http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
