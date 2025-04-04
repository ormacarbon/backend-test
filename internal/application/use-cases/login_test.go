package usecases_test

import (
	"errors"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*mocks.MockUserRepository, *mocks.MockTokenService, *usecases.LoginUserUseCase) {
	mockRepo := mocks.NewMockUserRepository()
	mockTokenService := mocks.NewMockTokenService()
	useCase := usecases.NewLoginUserUseCase(mockRepo, mockTokenService)
	return mockRepo, mockTokenService, useCase
}

func TestLoginUserUseCase_Execute(t *testing.T) {
	mockRepo, mockTokenService, useCase := setupTest()

	t.Run("success", func(t *testing.T) {

		password, _ := object_values.NewPassword("123456")
		email, _ := object_values.NewEmail("test@example.com")
		phoneNumber, _ := object_values.NewPhoneNumber("1234567890")

		user, _ := entities.NewUser("teste", email, password, phoneNumber)

		input := dto.LoginInput{Email: "test@example.com", Password: "123456"}
		expectedToken := "fake.jwt.token"

		mockRepo.On("FindByEmail", input.Email).Return(&user, nil)
		mockTokenService.On("GenerateToken", user.ID().String()).Return(expectedToken, nil)

		output, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, expectedToken, output.Token)
	})

	t.Run("user not found", func(t *testing.T) {
		input := dto.LoginInput{Email: "missing@example.com", Password: "any"}

		mockRepo.On("FindByEmail", input.Email).Return(nil, errors.New("not found"))

		_, err := useCase.Execute(input)
		assert.ErrorIs(t, err, shared.ErrNotFound)
	})

	t.Run("invalid password", func(t *testing.T) {
		email, _ := object_values.NewEmail("user@example.com")
		password, _ := object_values.NewPassword("correct-password")
		phone, _ := object_values.NewPhoneNumber("1234567890")

		user, _ := entities.NewUser("User", email, password, phone)

		input := dto.LoginInput{Email: "user@example.com", Password: "wrong-password"}

		mockRepo.On("FindByEmail", input.Email).Return(&user, nil)

		_, err := useCase.Execute(input)
		assert.ErrorIs(t, err, shared.ErrAuthorization)
	})

	t.Run("token generation error", func(t *testing.T) {
		email, _ := object_values.NewEmail("user2@example.com")
		password, _ := object_values.NewPassword("pass123")
		phone, _ := object_values.NewPhoneNumber("1234567890")

		user, _ := entities.NewUser("User", email, password, phone)

		input := dto.LoginInput{Email: "user2@example.com", Password: "pass123"}

		mockRepo.On("FindByEmail", input.Email).Return(&user, nil)
		mockTokenService.On("GenerateToken", user.ID().String()).Return("", errors.New("token error"))

		_, err := useCase.Execute(input)
		assert.ErrorIs(t, err, shared.ErrInternal)
	})
}
