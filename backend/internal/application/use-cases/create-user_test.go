package usecases

import (
	"strings"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest() (*mocks.MockUserRepository, *mocks.MockEmailService, *CreateUserUseCase) {
	mockRepo := mocks.NewMockUserRepository()
	mockEmail := mocks.NewMockEmailService()
	useCase := NewCreateUserUseCase(mockRepo, mockEmail)
	return mockRepo, mockEmail, useCase
}

func TestCreateUserUseCase_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo, _, useCase := setupTest()

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
		_, _, useCase := setupTest()

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
		_, _, useCase := setupTest()

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
		_, _, useCase := setupTest()

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
		mockRepo, _, useCase := setupTest()

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
		mockRepo, _, useCase := setupTest()

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

	t.Run("invited user", func(t *testing.T) {
		mockRepo, mockEmail, useCase := setupTest()

		validEmail, _ := object_values.NewEmail("inviter@example.com")
		validPhone, _ := object_values.NewPhoneNumber("+5511987654321")
		validPassword, _ := object_values.NewPassword("StrongP@ssw0rd")
		inviterCode := uuid.NewString()

		inviter := entities.LoadUser(
			uuid.New(),
			"Inviter",
			validEmail,
			validPassword,
			validPhone,
			inviterCode,
			nil,
			1,
		)

		input := dto.CreateUserInput{
			Name:       "New User",
			Email:      "newuser@example.com",
			Password:   "StrongP@ssw0rd",
			Phone:      "+5511999999999",
			InviteCode: &inviterCode,
		}

		mockRepo.On("FindByEmail", input.Email).Return(nil, nil)
		mockRepo.On("FindByInviteCode", *input.InviteCode).Return(&inviter, nil)
		mockRepo.On("Save", mock.Anything).Return(nil)
		mockEmail.On(
			"SendEmail",
			inviter.Email().Value(),
			"New user invited by you",
			mock.MatchedBy(func(body string) bool {
				return strings.Contains(body, "Hello Inviter") &&
					strings.Contains(body, "points") &&
					strings.Contains(body, "bvio")
			}),
		).Return(nil).Once()

		output, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.NotEmpty(t, output.UserID)
		mockRepo.AssertExpectations(t)
		mockEmail.AssertExpectations(t)
	})

	t.Run("should return error if invite code is invalid", func(t *testing.T) {
		mockRepo, mockEmail, useCase := setupTest()

		badCode := "invalid-uuid"

		input := dto.CreateUserInput{
			Name:       "Errored User",
			Email:      "errored@example.com",
			Password:   "StrongP@ssw0rd",
			Phone:      "+5511999999999",
			InviteCode: &badCode,
		}

		output, err := useCase.Execute(input)

		assert.Error(t, err)
		assert.Empty(t, output.UserID)
		mockRepo.AssertExpectations(t)
		mockEmail.AssertExpectations(t)
	})

	t.Run("should handle empty string invite code gracefully", func(t *testing.T) {
		mockRepo, mockEmail, useCase := setupTest()

		emptyCode := ""

		input := dto.CreateUserInput{
			Name:       "User Without Code",
			Email:      "userwithout@example.com",
			Password:   "StrongP@ssw0rd",
			Phone:      "+5511999999999",
			InviteCode: &emptyCode,
		}

		mockRepo.On("FindByEmail", input.Email).Return(nil, nil)
		mockRepo.On("Save", mock.Anything).Return(nil)

		output, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.NotEmpty(t, output.UserID)
		mockRepo.AssertExpectations(t)
		mockEmail.AssertExpectations(t)
	})

}
