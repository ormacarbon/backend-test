package service

import (
	"context"

	"github.com/icl00ud/backend-test/internal/model"
	"github.com/icl00ud/backend-test/internal/repository"
	"go.uber.org/zap"
)

type LeaderboardService interface {
	GetLeaderboard(ctx context.Context, limit int) ([]model.User, error)
}

type leaderboardServiceImpl struct {
	userRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewLeaderboardService(userRepo repository.UserRepository, logger *zap.SugaredLogger) LeaderboardService {
	return &leaderboardServiceImpl{
		userRepo: userRepo,
		logger:   logger.Named("LeaderboardService"),
	}
}

func (s *leaderboardServiceImpl) GetLeaderboard(ctx context.Context, limit int) ([]model.User, error) {
	users, err := s.userRepo.GetTopUsers(ctx, limit)
	if err != nil {
		s.logger.Errorw("Error retrieving top users", "limit", limit, "error", err)
		return nil, err
	}
	return users, nil
}
