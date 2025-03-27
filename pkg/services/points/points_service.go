package services

import (
	"gss-backend/pkg/models"
	repositories "gss-backend/pkg/repositories/points"
)

func NewPointsService(repository repositories.IPointsRepository) *PointsService {
	return &PointsService{
		repository: repository,
	}
}

func (s *PointsService) Create(user_id uint) (*models.Points, error) {
	return s.repository.Create(user_id)
}

func (s *PointsService) Update(user_id uint) (*models.Points, error) {
	return s.repository.Update(user_id)
}


func (s *PointsService) FindByUserID(user_id uint) (*models.Points, error) {
	return s.repository.FindByUserID(user_id)
}

func (s *PointsService) FindLeaderboard() (*[]models.Points, error) {
	return s.repository.CreateLeaderboard()
}