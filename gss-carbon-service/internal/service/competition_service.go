package service

import (
	"context"
	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/model"
	"github.com/icl00ud/backend-test/internal/repository"
)

type CompetitionService interface {
	FinishCompetition(ctx context.Context, limit int) ([]model.User, error)
}

type CompetitionServiceImpl struct {
	userRepo repository.UserRepository
	emailSvc email.EmailService
}

func NewCompetitionService(userRepo repository.UserRepository, emailSvc email.EmailService) CompetitionService {
	return &CompetitionServiceImpl{
		userRepo: userRepo,
		emailSvc: emailSvc,
	}
}

func (s *CompetitionServiceImpl) FinishCompetition(ctx context.Context, limit int) ([]model.User, error) {
	winners, err := s.userRepo.GetTopUsers(ctx, limit)
	if err != nil {
		return nil, err
	}

	subject := "Congratulations, you're a winner!"
	for _, user := range winners {
		body := "You have been selected as one of the top winners in the competition."
		_ = s.emailSvc.SendEmail(user.Email, subject, body)
	}

	return winners, nil
}
