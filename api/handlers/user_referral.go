package handlers

import (
	"gss-backend/api/presenters"
	services "gss-backend/pkg/services/user_referral"

	"github.com/gofiber/fiber/v2"
)


func FindLeaderboardScores(userReferralService services.IUserReferralService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := userReferralService.FindLeaderboardScores()

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.UserReferralErrorResponse(err))
		}

		return c.JSON(presenters.LeaderboardScoreSuceessResponse(result))
	}
}

