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

type LeaderboardServiceImpl struct {
	userRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewLeaderboardService(userRepo repository.UserRepository, logger *zap.SugaredLogger) LeaderboardService {
	return &LeaderboardServiceImpl{
		userRepo: userRepo,
		logger:   logger.Named("LeaderboardService"),
	}
}

func (s *LeaderboardServiceImpl) GetLeaderboard(ctx context.Context, limit int) ([]model.User, error) {
	s.logger.Infow("Getting leaderboard", "limit", limit)

	users, err := s.userRepo.GetTopUsers(ctx, limit)
	if err != nil {
		return nil, err
	}

	s.logger.Infow("Leaderboard retrieved successfully", "limit", limit, "count", len(users))
	return users, nil
}
