package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/backend-test/src/config"
	"github.com/joaooliveira247/backend-test/src/db"
	"github.com/joaooliveira247/backend-test/src/routes"
)

func init() {
	config.LoadEnv()
}

func main() {

	gormDB, err := db.GetDBConnection()

	if err != nil {
		log.Fatalf("DATABASE Connection: %v", err)
	}

	if err := db.CreateTables(gormDB); err != nil {
		log.Fatalf("DATABASE Tables Creation: %v", err)
		return
	}

	api := gin.Default()
	routes.RegistryRoutes(api)

	if err := api.Run(":8000"); err != nil {
		log.Fatalf("API RUN: err")
	}
}
