package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

type IUserReferralRepository interface {
	Create(referrerId uint, referredId uint) (*models.UserReferral, error)
	FindLeaderboard() (*[]LeaderboardScore, error)
}

type PostgresUserReferralRepository struct {
	db *gorm.DB
}

type LeaderboardScore struct {
	ReferrerId uint
	ReferralsCount int
}