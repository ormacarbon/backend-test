package service

import (
	"context"
	"sync"

	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/model"
	"github.com/icl00ud/backend-test/internal/repository"
	"go.uber.org/zap"
)

type CompetitionService interface {
	FinishCompetition(ctx context.Context, limit int) ([]model.User, error)
}

type competitionService struct {
	userRepo repository.UserRepository
	emailSvc email.Service
	logger   *zap.SugaredLogger
}

func NewCompetitionService(userRepo repository.UserRepository, emailSvc email.Service, logger *zap.SugaredLogger) CompetitionService {
	return &competitionService{
		userRepo: userRepo,
		emailSvc: emailSvc,
		logger:   logger.Named("CompetitionService"),
	}
}

func (s *competitionService) FinishCompetition(ctx context.Context, limit int) ([]model.User, error) {
	winners, err := s.userRepo.GetTopUsers(ctx, limit)
	if err != nil {
		s.logger.Errorw("Error retrieving top users", "limit", limit, "error", err)
		return nil, err
	}

	subject := "Congratulations, you're a winner!"
	templatePath := "internal/email/templates/winner_notification.html"

	var wg sync.WaitGroup
	for _, winner := range winners {
		wg.Add(1)
		go func(u model.User) {
			defer wg.Done()
			data := struct {
				Name   string
				Points int
			}{
				Name:   u.Name,
				Points: u.Points,
			}
			
			s.emailSvc.Email(u.Email, subject, templatePath, data)
		}(winner)
	}
	wg.Wait()

	// Clear the user table after sending emails
	if err := s.userRepo.CleanTable(ctx); err != nil {
		s.logger.Errorw("Failed to clean user table", "error", err)
		return nil, err
	}

	return winners, nil
}
