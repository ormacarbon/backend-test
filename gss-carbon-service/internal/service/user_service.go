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
	RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) (*model.User, error)
}

type userServiceImpl struct {
	emailSvc email.Service
	userRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewUserService(userRepo repository.UserRepository, emailSvc email.Service, logger *zap.SugaredLogger) UserService {
	return &userServiceImpl{
		emailSvc: emailSvc,
		userRepo: userRepo,
		logger:   logger.Named("UserService"),
	}
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, id string) (*model.User, error) {
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

func (s *userServiceImpl) GetUserByReferralToken(ctx context.Context, token string) (*model.User, error) {
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

func (s *userServiceImpl) GetReferrals(ctx context.Context, offset, limit int) ([]dto.ReferralPopulatedResponse, error) {
	referredUsers, err := s.userRepo.GetReferrals(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var responses []dto.ReferralPopulatedResponse
	for _, u := range referredUsers {
		referrerID := fmt.Sprintf("%d", *u.ReferredBy)
		referrer, err := s.userRepo.GetUserByID(ctx, referrerID)
		if err != nil {
			s.logger.Errorw("Failed to get referrer details", "referrerID", referrerID, "error", err)
			continue
		}
		if referrer == nil {
			s.logger.Warnw("Referrer not found", "referrerID", referrerID)
			continue
		}

		response := dto.ReferralPopulatedResponse{
			ID: u.ReferralToken,
			Referrer: dto.ReducedUserResponse{
				ID:    referrer.ID,
				Name:  referrer.Name,
				Email: referrer.Email,
			},
			Referred: dto.ReducedUserResponse{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			},
			CreatedAt: u.CreatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *userServiceImpl) RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) (*model.User, error) {
	existing, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		s.logger.Warnw("Email already exists", "email", req.Email)
		return nil, errs.New(errs.ErrEmailAlreadyExists.Error(), 409, errs.ErrEmailAlreadyExists)
	}

	var referrer *model.User
	if req.ReferralToken != "" {
		referrer, err = s.userRepo.GetUserByReferralToken(ctx, req.ReferralToken)
		if err != nil || referrer == nil {
			s.logger.Warnw("Invalid referral token", "token", req.ReferralToken, "error", err)
			return nil, errors.New("invalid referral token")
		}
		s.logger.Infow("Referrer found", "referrerID", referrer.ID)
	}

	newUser := &model.User{
		Name:          req.Name,
		Email:         req.Email,
		Phone:         req.Phone,
		Points:        1,
		ReferralToken: uuid.New().String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if referrer != nil {
		newUser.ReferredBy = &referrer.ID
	}

	if err := s.userRepo.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	if referrer != nil {
		if err := s.userRepo.UpdateUserPoints(ctx, referrer.ID, 1); err != nil {
			s.logger.Errorw("Failed to update referrer points", "referrerID", referrer.ID, "error", err)
		} else {
			s.logger.Infow("Awarded point to referrer", "referrerID", referrer.ID, "newUserID", newUser.ID)
			subject := "You've received an extra point!"
			templatePath := "internal/email/templates/referral_bonus.html"
			emailData := struct {
				Name   string
				Points int
			}{
				Name:   referrer.Name,
				Points: referrer.Points + 1,
			}

			s.emailSvc.Email(referrer.Email, subject, templatePath, emailData)
		}
	}

	return newUser, nil
}
