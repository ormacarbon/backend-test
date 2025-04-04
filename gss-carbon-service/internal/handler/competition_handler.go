package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/service"
	"go.uber.org/zap"
)

type CompetitionHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewCompetitionHandler(userService service.UserService, logger *zap.Logger) *CompetitionHandler {
	return &CompetitionHandler{
		userService: userService,
		logger:      logger.Named("CompetitionHandler"),
	}
}

// FinishCompetition finish the competition, sends notifications to winners and returns the list of winners.
func (h *CompetitionHandler) FinishCompetition(c *fiber.Ctx) error {
	sugar := h.logger.Sugar()
	sugar.Info("Received request to finish competition")

	limitParam := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		sugar.Warnw("Invalid limit query parameter, using default", "param", limitParam, "default", 10, "error", err)
		limit = 10
	} else {
		sugar.Infow("Using limit parameter", "limit", limit)
	}

	winners, err := h.userService.FinishCompetition(context.Background(), limit)
	if err != nil {
		sugar.Errorw("Failed to finish competition in service", "limit", limit, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to finish competition"})
	}

	sugar.Infow("Competition finished successfully", "winnerCount", len(winners))
	return c.JSON(winners)
}
