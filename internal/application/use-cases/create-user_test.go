package usecases

import (
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest() (*mocks.MockUserRepository, *CreateUserUseCase) {
	mockRepo := mocks.NewMockUserRepository()
	useCase := NewCreateUserUseCase(mockRepo)
	return mockRepo, useCase
}

func TestCreateUserUseCase_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo, useCase := setupTest()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "StrongP@ssw0rd",
			Phone:    "+5511987654321",
		}

		mockRepo.On("FindByEmail", input.Email).Return(nil, nil)
		mockRepo.On("Save", mock.Anything).Return(nil)

		output, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.NotEmpty(t, output.UserID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid email", func(t *testing.T) {
		_, useCase := setupTest()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "invalid-email",
			Password: "StrongP@ssw0rd",
			Phone:    "+5511987654321",
		}

		_, err := useCase.Execute(input)

		assert.ErrorIs(t, err, shared.ErrValidation)
	})

	t.Run("invalid phone", func(t *testing.T) {
		_, useCase := setupTest()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "StrongP@ssw0rd",
			Phone:    "1234",
		}

		_, err := useCase.Execute(input)

		assert.ErrorIs(t, err, shared.ErrValidation)
	})

	t.Run("invalid password", func(t *testing.T) {
		_, useCase := setupTest()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "123",
			Phone:    "+5511987654321",
		}

		_, err := useCase.Execute(input)

		assert.ErrorIs(t, err, shared.ErrValidation)
	})

	t.Run("user already exists", func(t *testing.T) {
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

		existingUser, _ := entities.NewUser("John Doe", email, password, phone, nil)

		mockRepo.On("FindByEmail", input.Email).Return(&existingUser, nil)

		_, err := useCase.Execute(input)

		assert.ErrorIs(t, err, shared.ErrConflictError)
		mockRepo.AssertNotCalled(t, "Save")
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository save error", func(t *testing.T) {
		mockRepo, useCase := setupTest()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "StrongP@ssw0rd",
			Phone:    "+5511987654321",
		}

		mockRepo.On("FindByEmail", input.Email).Return(nil, nil)
		mockRepo.On("Save", mock.Anything).Return(shared.ErrInternal)

		_, err := useCase.Execute(input)

		assert.EqualError(t, err, shared.ErrInternal.Error())
		mockRepo.AssertExpectations(t)
	})
}
