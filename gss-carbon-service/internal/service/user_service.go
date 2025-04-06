package service

import (
	"context"
	"errors"
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
	RegisterUser(ctx context.Context, user *dto.RegisterUserRequest) (*model.User, error)
	RegisterUserWithReferral(ctx context.Context, user *dto.RegisterUserWithReferralRequest) (*model.User, error)
}

type UserServiceImpl struct {
	emailSvc email.EmailService
	userRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewUserService(userRepo repository.UserRepository, emailSvc email.EmailService, logger *zap.SugaredLogger) UserService {
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

func (s *UserServiceImpl) RegisterUser(ctx context.Context, user *dto.RegisterUserRequest) (*model.User, error) {
	existingUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		s.logger.Warnw("Email already exists", "email", user.Email)
		return nil, errs.New(errs.ErrEmailAlreadyExists.Error(), 409, errs.ErrEmailAlreadyExists)
	}

	userToCreate := &model.User{
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Points:        1,
		ReferralToken: uuid.New().String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.userRepo.CreateUser(ctx, userToCreate); err != nil {
		return nil, err
	}

	return userToCreate, nil
}

func (s *UserServiceImpl) RegisterUserWithReferral(ctx context.Context, user *dto.RegisterUserWithReferralRequest) (*model.User, error) {
	s.logger.Infow("Registering new user with referral", "email", user.Email, "name", user.Name, "referralToken", user.ReferralToken)

	referrer, err := s.userRepo.GetUserByReferralToken(ctx, user.ReferralToken)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			s.logger.Warnw("Invalid referral token provided (referrer not found)", "token", user.ReferralToken)
			return nil, errors.New("invalid referral token") // Specific error for handler
		}
		// Other DB errors logged in repo
		return nil, err // Propagate other errors
	}
	s.logger.Infow("Referrer found", "referrerID", referrer.ID, "referrerEmail", referrer.Email)

	referralToken := uuid.New().String()
	userToCreate := &model.User{
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Points:        1, // Start with 1 point
		ReferralToken: referralToken,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.userRepo.CreateUser(ctx, userToCreate); err != nil {
		return nil, err
	}
	s.logger.Infow("Referred user created successfully", "userID", userToCreate.ID, "email", userToCreate.Email)

	if err := s.userRepo.UpdateUserPoints(ctx, referrer.ID, 1); err != nil {
		s.logger.Errorw("Failed to update referrer points after successful referral (continuing registration)",
			"referrerID", referrer.ID,
			"referredUserID", userToCreate.ID,
			"error", err,
		)
	} else {
		s.logger.Infow("Awarded point to referrer", "referrerID", referrer.ID, "referredUserID", userToCreate.ID)

		subject := "You've received an extra point!"
		body := "Congratulations! A new user has registered using your referral link, and you've earned an extra point."
		if emailErr := s.emailSvc.SendEmail(referrer.Email, subject, body); emailErr != nil {
			s.logger.Warnw("Failed to send referral bonus email",
				"recipient", referrer.Email,
				"referrerID", referrer.ID,
				"error", emailErr,
			)
		} else {
			s.logger.Infow("Sent referral bonus email", "recipient", referrer.Email, "referrerID", referrer.ID)
		}
	}

	return userToCreate, nil
}
