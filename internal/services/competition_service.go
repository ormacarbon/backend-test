package services

import (
	"github.com/jpeccia/go-backend-test/internal/models"
	"github.com/jpeccia/go-backend-test/internal/repositories"
)

type CompetitionService interface {
	GetTopWinners(limit int) ([]models.User, error)
}

type competitionService struct {
	userRepo repositories.UserRepository
}

func NewCompetitionService(userRepo repositories.UserRepository) CompetitionService {
	return &competitionService{userRepo: userRepo}
}

func (s *competitionService) GetTopWinners(limit int) ([]models.User, error) {
	return s.userRepo.FindTopUsersByPoints(limit)
}