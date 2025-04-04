package usecases

import (
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
)

type LoadUserByTokenUseCase struct {
	userRepo     output_ports.UserRepository
	tokenService output_ports.TokenService
}

func NewLoadUserByTokenUseCase(
	userRepo output_ports.UserRepository,
	tokenService output_ports.TokenService,
) *LoadUserByTokenUseCase {
	return &LoadUserByTokenUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (uc *LoadUserByTokenUseCase) Execute(token string) (*entities.User, error) {
	userID, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
