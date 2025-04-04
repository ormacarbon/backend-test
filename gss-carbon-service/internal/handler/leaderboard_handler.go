package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/service"
	"go.uber.org/zap"
)

type LeaderboardHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewLeaderboardHandler(userService service.UserService, logger *zap.Logger) *LeaderboardHandler {
	return &LeaderboardHandler{
		userService: userService,
		logger:      logger.Named("LeaderboardHandler"),
	}
}

// GetLeaderboard returns the top N users (default 10)
func (h *LeaderboardHandler) GetLeaderboard(c *fiber.Ctx) error {
	sugar := h.logger.Sugar()
	sugar.Info("Received request to get leaderboard")

	limitParam := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		sugar.Warnw("Invalid limit query parameter, using default", "param", limitParam, "default", 10, "error", err)
		limit = 10
	} else {
		sugar.Infow("Using limit parameter", "limit", limit)
	}

	users, err := h.userService.GetLeaderboard(context.Background(), limit)
	if err != nil {
		sugar.Errorw("Failed to get leaderboard from service", "limit", limit, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to retrieve leaderboard"})
	}

	sugar.Infow("Leaderboard retrieved successfully", "count", len(users))
	return c.JSON(users)
}
