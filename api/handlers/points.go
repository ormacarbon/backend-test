package handlers

import (
	"gss-backend/api/presenter"
	services "gss-backend/pkg/services/points"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdatePoints(points_service services.IPointsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			UserID string `json:"user_id"`
		}

		if err := c.BodyParser(&requestBody); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.PointsErrorResponse(err))
		}

		userIdStr := requestBody.UserID
		userId, err := strconv.ParseUint(userIdStr, 10, 32) 

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.PointsErrorResponse(err))
		}

		result, err := points_service.Update(uint(userId))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.PointsErrorResponse(err))
		}

		return c.JSON(presenter.PointsSuccessResponse(result))
	}
}

func FindLeaderboard(points_service services.IPointsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := points_service.FindLeaderboard()

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.PointsErrorResponse(err))
		}

		return c.JSON(presenter.PointsLeaderboardSuccessResponse(result))
	}
}

