package service

import (
	"context"
	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/model"
	"github.com/icl00ud/backend-test/internal/repository"
	"go.uber.org/zap"
	"sync"
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
		return nil, err
	}

	subject := "Congratulations, you're a winner!"
	templatePath := "internal/email/winner_notification.html"

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

	if err := s.userRepo.CleanTable(); err != nil {
		s.logger.Errorw("Failed to clean user table after competition", "error", err)
		return nil, err
	}

	return winners, nil
}
