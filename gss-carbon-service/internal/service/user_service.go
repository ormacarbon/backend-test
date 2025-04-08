package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/icl00ud/backend-test/internal/dto"
	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/errs"
	"github.com/icl00ud/backend-test/internal/model"
	"github.com/icl00ud/backend-test/internal/repository"
	"go.uber.org/zap"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByReferralToken(ctx context.Context, token string) (*model.User, error)
	GetReferrals(ctx context.Context, offset, limit int) ([]dto.ReferralPopulatedResponse, error)
	RegisterUser(ctx context.Context, user *dto.RegisterUserRequest) (*model.User, error)
}

type UserServiceImpl struct {
	emailSvc email.Service
	userRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewUserService(userRepo repository.UserRepository, emailSvc email.Service, logger *zap.SugaredLogger) UserService {
	return &UserServiceImpl{
		emailSvc: emailSvc,
		userRepo: userRepo,
		logger:   logger.Named("UserService"),
	}
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		s.logger.Warnw("User not found", "id", id)
		return nil, errs.New(errs.ErrUserNotFound.Error(), 404, errs.ErrUserNotFound)
	}

	return user, nil
}

func (s *UserServiceImpl) GetUserByReferralToken(ctx context.Context, token string) (*model.User, error) {
	user, err := s.userRepo.GetUserByReferralToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if user == nil {
		s.logger.Warnw("User not found", "token", token)
		return nil, errs.New(errs.ErrUserNotFound.Error(), 404, errs.ErrUserNotFound)
	}

	return user, nil
}

func (s *UserServiceImpl) GetReferrals(ctx context.Context, offset, limit int) ([]dto.ReferralPopulatedResponse, error) {
	referredUsers, err := s.userRepo.GetReferrals(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var responses []dto.ReferralPopulatedResponse
	for _, referredUser := range referredUsers {
		referrerID := fmt.Sprintf("%d", *referredUser.ReferredBy)
		referrer, err := s.userRepo.GetUserByID(ctx, referrerID)
		if err != nil {
			s.logger.Errorw("Failed to get referrer details", "referrerID", referrerID, "error", err)
			continue
		}

		response := dto.ReferralPopulatedResponse{
			ID: referredUser.ReferralToken,
			Referrer: dto.ReducedUserResponse{
				ID:    referrer.ID,
				Name:  referrer.Name,
				Email: referrer.Email,
			},
			Referred: dto.ReducedUserResponse{
				ID:    referredUser.ID,
				Name:  referredUser.Name,
				Email: referredUser.Email,
			},
			CreatedAt: referredUser.CreatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, user *dto.RegisterUserRequest) (*model.User, error) {
	existingUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		s.logger.Warnw("Email already exists", "email", user.Email)
		return nil, errs.New(errs.ErrEmailAlreadyExists.Error(), 409, errs.ErrEmailAlreadyExists)
	}

	var referrer *model.User
	if user.ReferralToken != "" {
		referrer, err = s.userRepo.GetUserByReferralToken(ctx, user.ReferralToken)
		if err != nil || referrer == nil {
			s.logger.Warnw("Invalid referral token provided", "token", user.ReferralToken, "error", err)
			return nil, errors.New("invalid referral token")
		}
		s.logger.Infow("Referrer found", "referrerID", referrer.ID, "referrerEmail", referrer.Email)
	}

	newUser := &model.User{
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Points:        1, // Starts with 1 point
		ReferralToken: uuid.New().String(),
		ReferredBy:    nil,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if referrer != nil {
		newUser.ReferredBy = &referrer.ID
	}

	if err := s.userRepo.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	// If the user was referred, update the referrer points and email
	if referrer != nil {
		if err := s.userRepo.UpdateUserPoints(ctx, referrer.ID, 1); err != nil {
			s.logger.Errorw("Failed to update referrer points", "referrerID", referrer.ID, "error", err)
		} else {
			s.logger.Infow("Awarded point to referrer", "referrerID", referrer.ID, "referredUserID", newUser.ID)

			subject := "You've received an extra point!"
			templatePath := "internal/email/templates/referral_bonus.html"
			emailData := struct {
				Name   string
				Points int
			}{
				Name:   newUser.Name,
				Points: referrer.Points + 1,
			}

			s.emailSvc.Email(referrer.Email, subject, templatePath, emailData)
		}
	}

	return newUser, nil
}
