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
	FullName   string `json:"full_name"`
	ReferralsCount uint `json:"referrals_count"`
}

func LeaderboardScoreSuceessResponse(data *[]repo.LeaderboardScore) *fiber.Map {
	var leaderboardScores []LeaderboardScorePresenter

	for _, score := range *data {
		leaderboardScores = append(leaderboardScores, LeaderboardScorePresenter{
			ReferrerId: score.ReferrerId,
			FullName:   score.FullName,
			ReferralsCount: score.ReferralsCount,
		})
	}
	return &fiber.Map{
		"status": "success",
		"data":   leaderboardScores,
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