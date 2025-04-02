package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-backend-test/config"
	"github.com/jpeccia/go-backend-test/internal/database/migrations"
	"github.com/jpeccia/go-backend-test/internal/handlers"
	"github.com/jpeccia/go-backend-test/internal/repositories"
	"github.com/jpeccia/go-backend-test/internal/routes"
	"github.com/jpeccia/go-backend-test/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()
	migrations.Migrate()

	r := gin.Default()

	userRepo := repositories.NewUserRepository(config.DB)

	userService := services.NewUserService(userRepo)
	competitionService := services.NewCompetitionService(userRepo)

	userHandler := handlers.NewUserHandler(userService)
	competitionHandler := handlers.NewCompetitionHandler(competitionService)

	routes.SetupRoutes(r, userHandler, competitionHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
