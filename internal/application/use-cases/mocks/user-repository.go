package mocks

import (
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if user := args.Get(0); user != nil {
		return user.(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Save(user entities.User) error {
	args := m.Called(user)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func NewMockUserRepository() *MockUserRepository {
	return new(MockUserRepository)
}
