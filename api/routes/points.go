package routes

import (
	"gss-backend/api/handlers"
	services "gss-backend/pkg/services/points"

	"github.com/gofiber/fiber/v2"
)

func PointsRouter(app fiber.Router, points_service services.IPointsService) {
	app.Post("/points", handlers.UpdatePoints(points_service))
	app.Get("/points/leaderboard", handlers.FindLeaderboard(points_service))
}