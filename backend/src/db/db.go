package db

import (
	"github.com/joaooliveira247/backend-test/src/config"
	"github.com/joaooliveira247/backend-test/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(config.DATABASE_URL),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Users{}, &models.Competitions{}, &models.Points{}); err != nil {
		return err
	}
	return nil
}
