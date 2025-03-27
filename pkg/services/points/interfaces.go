package services

import (
	"gss-backend/pkg/models"
	repositories "gss-backend/pkg/repositories/points"
)

type IPointsService interface {
	Create(user_id uint) (*models.Points, error)
	Update(user_id uint) (*models.Points, error)
	FindByUserID(user_id uint) (*models.Points, error)
	FindLeaderboard() (*[]models.Points, error)
}

type PointsService struct {
	repository repositories.IPointsRepository
}