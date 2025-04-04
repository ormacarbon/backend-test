package repositories

import "github.com/cassiusbessa/backend-test/internal/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
