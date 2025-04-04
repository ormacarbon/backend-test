package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthHandler struct {
	DB     *gorm.DB
	logger *zap.Logger
}

func NewHealthHandler(db *gorm.DB, logger *zap.Logger) *HealthHandler {
	return &HealthHandler{
		DB:     db,
		logger: logger.Named("HealthHandler"),
	}
}

func (h *HealthHandler) Ping(c *fiber.Ctx) error {
	h.logger.Debug("Received health ping request")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (h *HealthHandler) Checker(c *fiber.Ctx) error {
	sugar := h.logger.Sugar()
	sugar.Info("Received health check request")

	sqlDB, err := h.DB.DB()
	if err != nil {
		sugar.Errorw("Failed to get underlying *sql.DB from GORM", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"error":  "database connection error",
		})
	}

	if err = sqlDB.Ping(); err != nil {
		sugar.Errorw("Database ping failed", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"error":  "database ping failed",
		})
	}

	sugar.Info("Health check successful: Database connected")
	return c.JSON(fiber.Map{
		"status":   "ok",
		"database": "connected",
	})
}
