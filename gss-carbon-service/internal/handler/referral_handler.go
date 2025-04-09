package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/service"
	"go.uber.org/zap"
)

type ReferralsHandler struct {
	userService service.UserService
	logger      *zap.SugaredLogger
}

func NewReferralsHandler(userService service.UserService, logger *zap.SugaredLogger) *ReferralsHandler {
	return &ReferralsHandler{
		userService: userService,
		logger:      logger.Named("ReferralsHandler"),
	}
}

// GetReferrals returns a paginated list of referrals
func (h *ReferralsHandler) GetReferrals(c *fiber.Ctx) error {
	h.logger.Info("Get referrals request received")

	pageParam := c.Query("page", "1")
	pageSizeParam := c.Query("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	referrals, err := h.userService.GetReferrals(context.Background(), offset, pageSize)
	if err != nil {
		h.logger.Errorw("Failed to get referrals", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get referrals"})
	}

	if len(referrals) == 0 {
		h.logger.Warn("No referrals found")
	}

	return c.Status(fiber.StatusOK).JSON(referrals)
}
