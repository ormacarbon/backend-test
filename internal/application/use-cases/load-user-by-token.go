package usecases

import (
	"github.com/cassiusbessa/backend-test/internal/application/dto"
	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
)

type LoadUserByTokenUseCase struct {
	userRepo     output_ports.UserRepository
	tokenService output_ports.TokenService
}

func NewLoadUserByTokenUseCase(
	userRepo output_ports.UserRepository,
	tokenService output_ports.TokenService,
) input_ports.LoadUserByTokenUseCase {
	return LoadUserByTokenUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (uc LoadUserByTokenUseCase) Execute(token string) (*dto.LoadedUserOutput, error) {
	userID, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}
	return &dto.LoadedUserOutput{
		ID:    user.ID().String(),
		Name:  user.Name(),
		Email: user.Email().Value(),
		Phone: user.Phone().Value(),
	}, nil
}
