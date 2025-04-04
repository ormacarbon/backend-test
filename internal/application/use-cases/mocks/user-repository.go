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

func (m *MockUserRepository) FindByID(id string) (*entities.User, error) {
	args := m.Called(id)
	if user := args.Get(0); user != nil {
		return user.(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByInviteCode(inviteCode string) (*entities.User, error) {
	args := m.Called(inviteCode)
	if user := args.Get(0); user != nil {
		return user.(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func NewMockUserRepository() *MockUserRepository {
	return new(MockUserRepository)
}
