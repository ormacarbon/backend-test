package presenter

import (
	"gss-backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type PointsPresenter struct {
	ID        uint   `json:"id"`
	Points   uint   `json:"points"`
}

func PointsSuccessResponse(data *models.Points) *fiber.Map {
	points := PointsPresenter{
		ID:        data.ID,
		Points:    data.Points,
	}

	return &fiber.Map{
		"status": "success",
		"data":   points,
		"error":  nil,
	}
}

func PointsLeaderboardSuccessResponse(data *[]models.Points) *fiber.Map {
	var points []PointsPresenter

	for _, point := range *data {
		points = append(points, PointsPresenter{
			ID:        point.ID,
			Points:    point.Points,
		})
	}

	return &fiber.Map{
		"status": "success",
		"data":   points,
		"error":  nil,
	}
}

func PointsErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"data":   nil,
		"error":  err.Error(),
	}
}