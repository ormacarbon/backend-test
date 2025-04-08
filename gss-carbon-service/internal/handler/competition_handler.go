package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/service"
	"go.uber.org/zap"
)

type CompetitionHandler struct {
	competitionService service.CompetitionService
	logger             *zap.SugaredLogger
}

func NewCompetitionHandler(competitionService service.CompetitionService, logger *zap.SugaredLogger) *CompetitionHandler {
	return &CompetitionHandler{
		competitionService: competitionService,
		logger:             logger.Named("CompetitionHandler"),
	}
}

// FinishCompetition ends the competition, notifies winners, and returns the list of winners
func (h *CompetitionHandler) FinishCompetition(c *fiber.Ctx) error {
	h.logger.Info("Received request to finish competition")

	limitParam := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		h.logger.Warnw("Invalid limit query parameter, using default", "param", limitParam, "default", 10, "error", err)
		limit = 10
	} else {
		h.logger.Infow("Using limit parameter", "limit", limit)
	}

	winners, err := h.competitionService.FinishCompetition(context.Background(), limit)
	if err != nil {
		h.logger.Errorw("Failed to finish competition", "limit", limit, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to finish competition"})
	}

	h.logger.Infow("Competition finished successfully", "winnerCount", len(winners))
	return c.Status(fiber.StatusOK).JSON(winners)
}
