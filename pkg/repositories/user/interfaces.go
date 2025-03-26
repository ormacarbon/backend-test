package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() (*[]models.User, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)	
}

type PostgresUserRepository struct {
	db *gorm.DB
}