package input_ports

import "github.com/cassiusbessa/backend-test/internal/domain/entities"

type LoadUserByTokenUseCase interface {
	Execute(token string) (entities.User, error)
}
