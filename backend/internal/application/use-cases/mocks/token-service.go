package mocks

import "github.com/stretchr/testify/mock"

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateToken(userID string) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) ValidateToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func NewMockTokenService() *MockTokenService {
	return new(MockTokenService)
}
