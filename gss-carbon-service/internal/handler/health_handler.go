package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthHandler struct {
	DB     *gorm.DB
	logger *zap.SugaredLogger
}

func NewHealthHandler(db *gorm.DB, logger *zap.SugaredLogger) *HealthHandler {
	return &HealthHandler{
		DB:     db,
		logger: logger.Named("HealthHandler"),
	}
}

// Ping returns a simple status message
func (h *HealthHandler) Ping(c *fiber.Ctx) error {
	h.logger.Debug("Received health ping request")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

// Checker verifies if the database connection is healthy
func (h *HealthHandler) Checker(c *fiber.Ctx) error {
	h.logger.Info("Received health check request")

	sqlDB, err := h.DB.DB()
	if err != nil {
		h.logger.Errorw("Failed to obtain *sql.DB from GORM", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"error":  "database connection error",
		})
	}

	if err = sqlDB.Ping(); err != nil {
		h.logger.Errorw("Database ping failed", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"error":  "database ping failed",
		})
	}

	h.logger.Info("Health check successful: Database connected")
	return c.JSON(fiber.Map{
		"status":   "ok",
		"database": "connected",
	})
}
