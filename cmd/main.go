package main

import (
	"log"
	"os"

	"github.com/cassiusbessa/backend-test/internal/infra/db"
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

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Server running on port %s 🚀", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
