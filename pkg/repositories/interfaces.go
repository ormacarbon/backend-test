package repositories

import (
	"gss-backend/pkg/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (models.User, error)
	Create(user models.User) (models.User, error)	
}