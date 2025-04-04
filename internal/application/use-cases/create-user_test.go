package usecases

import (
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	user, _ := args.Get(0).(*entities.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) SetupDefaultBehavior() {
	m.On("FindByEmail", mock.Anything).Return(nil, nil)
	m.On("Save", mock.Anything).Return(nil)
}

func setupTest() (*MockUserRepository, *CreateUserUseCase) {
	mockRepo := new(MockUserRepository)
	mockRepo.SetupDefaultBehavior()
	useCase := NewCreateUserUseCase(mockRepo)
	return mockRepo, useCase
}

func TestCreateUserUseCase_Execute_Success(t *testing.T) {
	mockRepo, useCase := setupTest()

	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "StrongP@ssw0rd",
		Phone:    "+5511987654321",
	}

	output, err := useCase.Execute(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, output.UserID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUserUseCase_Execute_InvalidEmail(t *testing.T) {
	_, useCase := setupTest()

	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "invalid-email",
		Password: "StrongP@ssw0rd",
		Phone:    "+5511987654321",
	}

	_, err := useCase.Execute(input)

	assert.ErrorIs(t, err, shared.ErrValidation)
}

func TestCreateUserUseCase_Execute_InvalidPhone(t *testing.T) {
	_, useCase := setupTest()

	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "StrongP@ssw0rd",
		Phone:    "1234",
	}

	_, err := useCase.Execute(input)

	assert.ErrorIs(t, err, shared.ErrValidation)
}

func TestCreateUserUseCase_Execute_InvalidPassword(t *testing.T) {
	_, useCase := setupTest()

	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "123",
		Phone:    "+5511987654321",
	}

	_, err := useCase.Execute(input)

	assert.ErrorIs(t, err, shared.ErrValidation)
}

func TestCreateUserUseCase_Execute_SaveError(t *testing.T) {
	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "StrongP@ssw0rd",
		Phone:    "+5511987654321",
	}
	mockRepo, useCase := setupTest()

	mockRepo.ExpectedCalls = nil
	mockRepo.On("Save", mock.Anything).Return(shared.ErrInternal)
	mockRepo.On("FindByEmail", input.Email).Return(nil, nil)

	_, err := useCase.Execute(input)

	mockRepo.AssertCalled(t, "Save", mock.Anything)
	assert.EqualError(t, err, shared.ErrInternal.Error())
	mockRepo.AssertExpectations(t)
}

func TestCreateUserUseCase_Execute_UserAlreadyExists(t *testing.T) {
	mockRepo, useCase := setupTest()

	input := dto.CreateUserInput{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "StrongP@ssw0rd",
		Phone:    "+5511987654321",
	}

	email, _ := object_values.NewEmail(input.Email)
	phone, _ := object_values.NewPhoneNumber(input.Phone)
	password, _ := object_values.NewPassword(input.Password)

	existingUser, _ := entities.NewUser(
		"John Doe",
		email,
		password,
		phone,
	)

	mockRepo.ExpectedCalls = nil
	mockRepo.On("FindByEmail", input.Email).Return(&existingUser, nil)

	_, err := useCase.Execute(input)

	assert.ErrorIs(t, err, shared.ErrConflictError)
	mockRepo.AssertNotCalled(t, "Save")
	mockRepo.AssertExpectations(t)
}
