package services

import (
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
)

type IUserReferralService interface {
	FindLeaderboardScores() (*[]userReferralRepo.LeaderboardScore, error)
}

type UserReferralService struct {
	userReferralRepo userReferralRepo.IUserReferralRepository
	userRepo userRepo.IUserRepository
}