package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jpeccia/go-backend-test/internal/dto"
	"github.com/jpeccia/go-backend-test/internal/models"
	"github.com/jpeccia/go-backend-test/internal/repositories"
)

type UserService interface {
	RegisterUser(dto dto.RegisterUserDTO) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) RegisterUser(dto dto.RegisterUserDTO) (*models.User, error) {
	referralCode := uuid.New().String()

	user := models.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PhoneNumber:  dto.PhoneNumber,
		ReferralCode: referralCode,
		Points:       1,
	}

	emailService := NewEmailService()

	if dto.ReferredBy != "" {
		referredUser, err := u.userRepo.FindUserByReferralCode(dto.ReferredBy)
		if err != nil {
			return nil, errors.New("invalid referral code")
		}

		referredUser.Points++
		err = u.userRepo.UpdateUser(referredUser)
		if err != nil {
			return nil, err
		}

		user.ReferredBy = referredUser.Email

		go func() {
			err := emailService.SendEmail(
				referredUser.Email,
				"Congratulations! You earned an extra point!",
				fmt.Sprintf("A new person signed up using your referral link! You now have %d points.", referredUser.Points),
			)
			if err != nil {
				fmt.Println("Erro ao enviar e-mail para o referenciador:", err)
			}
		}()
	}

	err := u.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	go func() {
		err := emailService.SendEmail(
			user.Email,
			"Welcome to the competition!",
			fmt.Sprintf("Hi %s :), you have successfully registered! Share your link to earn points: %s/signup?ref=%s",
				user.Name, os.Getenv("FRONTEND_URL"), user.ReferralCode),
		)
		if err != nil {
			fmt.Println("Erro ao enviar e-mail de boas-vindas:", err)
		}
	}()

	return &user, nil
}
