package services

import (
	"gss-backend/api/dtos"
	"gss-backend/pkg/models"
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
	emailService "gss-backend/pkg/services/email"
)

type IUserService interface {
	FindAll() (*[]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByReferralCode(referralCode string) (*models.User, error)
	Create(userDto *dtos.CreateUserDTO) (*models.User, error)
}

type UserService struct {
	userRepo userRepo.IUserRepository
	userReferralRepo userReferralRepo.IUserReferralRepository
	emailService emailService.IEmailService
}