package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

// Concrete implementation of the IPointsRepository interface
func NewPostgresUserReferralRepository(db *gorm.DB) *PostgresUserReferralRepository {
	return &PostgresUserReferralRepository{db: db}
}

// Create a new user referral
func (r *PostgresUserReferralRepository) Create(referrerId uint, referredId uint) (*models.UserReferral, error) {
	userReferral := models.UserReferral{ReferrerId: referrerId, ReferredId: referredId}
	result := r.db.Create(&userReferral)
	return &userReferral, result.Error
	
}

// Find the top 10 users with the most referrals
func (r *PostgresUserReferralRepository) FindLeaderboard() (*[]LeaderboardScore, error) {
	var leaderboardScores []LeaderboardScore

	result := r.db.Table("user_referrals").
        Select("referrer_id, COUNT(*) as referrals_count").
        Group("referrer_id").
        Order("referrals_count DESC").
        Limit(10).
        Find(&leaderboardScores)


	return &leaderboardScores, result.Error
}

