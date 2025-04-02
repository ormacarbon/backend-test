package services

import (
	"fmt"
	"log"

	"github.com/jpeccia/go-backend-test/internal/models"
	"github.com/jpeccia/go-backend-test/internal/repositories"
)

type CompetitionService interface {
	GetTopWinners(limit int) ([]models.User, error)
	NotifyWinners(winners []models.User) error
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

func (s *competitionService) NotifyWinners(winners []models.User) error {
	emailService := NewEmailService()

	for _, winner := range winners {
		subject := "Congratulations! You are a winner!"
		body := fmt.Sprintf("Dear %s,\n\nYou are one of the top 10 winners of the competition! Congratulations!\n\nBest regards,\nCompetition Team", winner.Name)

		err := emailService.SendEmail(winner.Email, subject, body)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", winner.Email, err)
		}
	}

	return nil
}