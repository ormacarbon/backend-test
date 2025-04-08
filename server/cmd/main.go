package main

import (
	"log"
	"server/api/handlers"
	"server/api/routes"
	"server/internal/config"
	"server/internal/controllers"
	"server/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	userRepo := repository.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepo)
	userHandler := handlers.NewUserHandler(userController)

	routes.SetupRoutes(r, userHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
