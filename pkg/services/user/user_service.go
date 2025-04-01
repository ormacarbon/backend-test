package services

import (
	"fmt"
	"gss-backend/api/dtos"
	"gss-backend/pkg/models"
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
	emailService "gss-backend/pkg/services/email"
	"gss-backend/pkg/utils"
)

// Instatiate a new UserService
func NewUserService(
	userRepo userRepo.IUserRepository,
	userReferralRepo userReferralRepo.IUserReferralRepository,
	emailService emailService.IEmailService ) *UserService {
	return &UserService{
		userRepo: userRepo,
		userReferralRepo: userReferralRepo,
		emailService: emailService,

	}
}

// Register a new user and create a new points record for the user
func (s *UserService) Create(userDto *dtos.CreateUserDTO) (*models.User, error) {
	// Generate new referral code for the user
	referralCode := utils.GenerateReferralCode()


	// Create a new user model
	newUser := models.User{
		FullName: userDto.FullName,
		Email: userDto.Email,
		PhoneNumber: userDto.PhoneNumber,
		ReferralCode: referralCode,
	}

	// Create a new user record in the database
	createdUser, err := s.userRepo.Create(&newUser)

	if err != nil {
		return nil, err
	}

	// Send welcome email to the user asynchronously
	go func() {
		err := s.emailService.SendWelcomeEmail(createdUser.Email)

		if err != nil {
			fmt.Printf("Failed to send welcome email to %s: %v\n", createdUser.Email, err)
		}
	}()
	

	// Create user referral record for the registered user
	_, err = s.userReferralRepo.Create(createdUser.ID, createdUser.ID)

	if err != nil {
		return nil, err
	}

	// If user has a referral code, set another record for the referrer
	if userDto.ReferrerCode != "" {
		referrerUser, err := s.userRepo.FindByReferralCode(userDto.ReferrerCode)

		if err != nil {
			return nil, err
		}

		_, err = s.userReferralRepo.Create(referrerUser.ID, createdUser.ID)

		if err != nil {
			return nil, err
		}

		// Send email to the referrer user asynchronously
		go func() {
			// Rewrote in another way so I can remind this also works hehe
			if err := s.emailService.SendReferralLinkAccess(referrerUser.Email); err != nil {
				fmt.Printf("Error sending referral link access email to %s: %v\n", referrerUser.Email, err)
			}
		}()
	}

	return createdUser, nil
}

// Find all users (used for developing purposes)
func (s *UserService) FindAll() (*[]models.User, error) {
	return s.userRepo.FindAll()
}

// Find a user by their ID (used for developing purposes)
func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

// Find a user by their referral code
func (s *UserService) FindByReferralCode(referralCode string) (*models.User, error) {
	return s.userRepo.FindByReferralCode(referralCode)
}





