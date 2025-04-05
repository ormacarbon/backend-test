package usecases

import (
	"github.com/cassiusbessa/backend-test/internal/application/dto"
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
)

type LoginUserUseCase struct {
	userRepo     output_ports.UserRepository
	tokenService output_ports.TokenService
}

func NewLoginUserUseCase(
	userRepo output_ports.UserRepository,
	tokenService output_ports.TokenService,
) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (uc *LoginUserUseCase) Execute(input dto.LoginInput) (dto.LoginOutput, error) {
	user, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return dto.LoginOutput{}, shared.ErrInternal
	}

	if user == nil {
		return dto.LoginOutput{}, shared.ErrAuthorization
	}

	if !user.Password().Compare(input.Password) {
		return dto.LoginOutput{}, shared.ErrAuthorization
	}

	token, err := uc.tokenService.GenerateToken(user.ID().String())
	if err != nil {
		return dto.LoginOutput{}, shared.ErrInternal
	}

	return dto.LoginOutput{Token: token}, nil
}
