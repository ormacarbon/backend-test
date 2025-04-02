package services

import (
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
	emailService "gss-backend/pkg/services/email"
)

type IUserReferralService interface {
	FindLeaderboardScores() (*[]userReferralRepo.LeaderboardScore, error)
}

type UserReferralService struct {
	userReferralRepo userReferralRepo.IUserReferralRepository
	userRepo userRepo.IUserRepository
	emailService emailService.IEmailService
}