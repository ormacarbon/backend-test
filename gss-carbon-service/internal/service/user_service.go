package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(ctx context.Context, name, email, phone string) (*model.User, error)
	RegisterUserWithReferral(ctx context.Context, name, email, phone, referralToken string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, name, email, phone string) (*model.User, error) {
	token := uuid.New().String()
	user := &model.User{
		Name:          name,
		Email:         email,
		Phone:         phone,
		Points:        1, // Ponto inicial
		ReferralToken: token,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) RegisterUserWithReferral(ctx context.Context, name, email, phone, referralToken string) (*model.User, error) {
	// Busca o usuário que gerou o referral
	referrer, err := s.userRepo.GetUserByReferralToken(ctx, referralToken)
	if err != nil {
		return nil, errors.New("invalid referral token")
	}

	token := uuid.New().String()
	user := &model.User{
		Name:          name,
		Email:         email,
		Phone:         phone,
		Points:        1,
		ReferralToken: token,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	// Incrementa os pontos do usuário que fez o referral
	if err := s.userRepo.UpdateUserPoints(ctx, referrer.ID, 1); err != nil {
		return nil, err
	}

	// Aqui pode ser chamado o módulo de e-mail para notificar o referrer
	return user, nil
}
