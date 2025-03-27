package services

import (
	"gss-backend/pkg/models"
	pointsRepo "gss-backend/pkg/repositories/points"
	userRepo "gss-backend/pkg/repositories/user"
)

type IUserService interface {
	FindAll() (*[]models.User, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type UserService struct {
	pointsRepo pointsRepo.IPointsRepository
	userRepo userRepo.IUserRepository
}