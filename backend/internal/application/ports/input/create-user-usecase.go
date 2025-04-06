package input_ports

import "github.com/cassiusbessa/backend-test/internal/application/dto"

type CreateUserUseCase interface {
	Execute(input dto.CreateUserInput) (dto.CreateUserOutput, error)
}
