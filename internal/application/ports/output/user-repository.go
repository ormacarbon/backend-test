package output_ports

import "github.com/cassiusbessa/backend-test/internal/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
}
