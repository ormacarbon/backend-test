package services

import (
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
)



func NewUserReferralService(userRepo userRepo.IUserRepository, userReferralRepo userReferralRepo.IUserReferralRepository) IUserReferralService {
	return &UserReferralService{
		userRepo: userRepo,
		userReferralRepo: userReferralRepo,
	}
}

// Function that calculates the appropiate score for the leaderboard
func (s *UserReferralService) FindLeaderboardScores() (*[]userReferralRepo.LeaderboardScore, error) {
	// Find the top 10 users with the most referrals
	leaderboardScores, err := s.userReferralRepo.FindLeaderboard()

	if err != nil {
		return nil, err
	}

	// Add 1 point to each Referral count
	for _, score := range *leaderboardScores {
		score.ReferralsCount += 1
	}

	return leaderboardScores, nil
}
