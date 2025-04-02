package services

import (
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

	if dto.ReferredBy != "" {
		referredUser, err := u.userRepo.FindUserByEmail(dto.ReferredBy)
		if err == nil {
			referredUser.Points++
			u.userRepo.CreateUser(referredUser)
		}
	}

	err := u.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

