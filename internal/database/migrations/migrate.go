package migrations

import (
	"log"

	"github.com/jpeccia/go-backend-test/config"
	"github.com/jpeccia/go-backend-test/internal/models"
)

func Migrate() {
	db := config.DB

	log.Println("Running database migrations...")

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrated successfully!")
}