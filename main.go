package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpeccia/go-backend-test/config"
	"github.com/jpeccia/go-backend-test/internal/database/migrations"
	"github.com/jpeccia/go-backend-test/internal/handlers"
	"github.com/jpeccia/go-backend-test/internal/repositories"
	"github.com/jpeccia/go-backend-test/internal/routes"
	"github.com/jpeccia/go-backend-test/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	migrations.Migrate()
	
	r := gin.Default()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes.SetupRoutes(r, userHandler)

	r.Run(":8080")
}