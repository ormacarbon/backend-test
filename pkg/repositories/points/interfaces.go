package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

type IPointsRepository interface {
	Create(user_id uint) (*models.Points, error)
	Update(user_id uint) (*models.Points, error)
	FindByUserID(user_id uint) (*models.Points, error)
	CreateLeaderboard() (*[]models.Points, error)
}

type PostgresPointsRepository struct {
	db *gorm.DB
}