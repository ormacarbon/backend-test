package routes

import (
	"gss-backend/api/handlers"
	services "gss-backend/pkg/services/user_referral"

	"github.com/gofiber/fiber/v2"
)


func UserReferralRouter(app fiber.Router, user_referral_service services.IUserReferralService) {
	app.Get("/leaderboard", handlers.FindLeaderboardScores(user_referral_service))
}