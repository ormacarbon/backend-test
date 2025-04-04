package input_ports

import (
	"github.com/cassiusbessa/backend-test/internal/application/dto"
)

type LoadUserByTokenUseCase interface {
	Execute(token string) (*dto.LoadedUserOutput, error)
}
