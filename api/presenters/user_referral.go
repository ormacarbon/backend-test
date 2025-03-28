package presenters

import (
	repo "gss-backend/pkg/repositories/user_referral"

	"github.com/gofiber/fiber/v2"
)

type PointsPresenter struct {
	ID        uint   `json:"id"`
	Points   uint   `json:"points"`
}

type LeaderboardScorePresenter struct {
	ReferrerId uint   `json:"referrer_id"`
	ReferralsCount uint `json:"referrals_count"`
}

func LeaderboardScoreSuceessResponse(data *[]repo.LeaderboardScore) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func UserReferralErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"data":   nil,
		"error":  err.Error(),
	}
}