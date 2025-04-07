package input_ports

import "github.com/cassiusbessa/backend-test/internal/application/dto"

type LoginUseCase interface {
	Execute(input dto.LoginInput) (dto.LoginOutput, error)
}
