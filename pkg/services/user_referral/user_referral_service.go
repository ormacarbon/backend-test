package services

import (
	"fmt"
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
	emailService "gss-backend/pkg/services/email"
)



func NewUserReferralService(userRepo userRepo.IUserRepository,
	userReferralRepo userReferralRepo.IUserReferralRepository,
	emailService emailService.IEmailService) IUserReferralService {
	return &UserReferralService{
		userRepo: userRepo,
		userReferralRepo: userReferralRepo,
		emailService: emailService,
	}
}

// Function that calculates the appropiate score for the leaderboard
func (s *UserReferralService) FindLeaderboardScores() (*[]userReferralRepo.LeaderboardScore, error) {
	// Find the top 10 users with the most referrals
	leaderboardScores, err := s.userReferralRepo.FindLeaderboard()

	if err != nil {
		return nil, err
	}

	// Instatiating error channel
	errChan := make(chan error, len(*leaderboardScores))

	// For every user in the leaderboard, send an email to say he is in the leaderboard
	for _, leaderboardScore := range *leaderboardScores {
		// Get user data and send email asynchronously
		// If any step causes an error, send the error to the error channel
		go func(leaderboardScore userReferralRepo.LeaderboardScore) {
			user, err := s.userRepo.FindByID(leaderboardScore.ReferrerId)

			if err != nil {
				errChan <- fmt.Errorf("Failed to find user with ID %d: %v", leaderboardScore.ReferrerId, err)
				return
			}

			err = s.emailService.SendLeaderboardEmail(user.Email)

			if err != nil {
				errChan <- fmt.Errorf("Failed to send leaderboard email to %s: %v", user.Email, err)
				return
			}

			errChan <- nil // Signal that the email was sent successfully
		}(leaderboardScore)

	}

	// Processing errors in another goroutine
	go func() {
		for range *leaderboardScores {
			if err := <-errChan; err != nil {
				fmt.Println("Error processing leaderboard email: ", err)
			}
		}

		close(errChan)
	}()

	return leaderboardScores, nil
}
