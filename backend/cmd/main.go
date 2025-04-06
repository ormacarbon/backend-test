package main

import (
	"log"
	"os"
	"time"

	"github.com/cassiusbessa/backend-test/internal/infra/db"
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	routes.WithCreateUser(api)
	routes.WithLogin(api)
	routes.WithLoadUserByToken(api)
	routes.WithUsersRanking(api)

	log.Printf("Server running on port %s 🚀", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
