package usecases_test

import (
	"testing"

	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func validEmail() object_values.Email {
	email, _ := object_values.NewEmail("john@example.com")
	return email
}

func validPhone() object_values.PhoneNumber {
	phone, _ := object_values.NewPhoneNumber("11999999999")
	return phone
}

func validPassword() object_values.Password {
	password, _ := object_values.NewPassword("123456")
	return password
}

func setupLoadUserByTokenTest() (*mocks.MockUserRepository, *mocks.MockTokenService, input_ports.LoadUserByTokenUseCase) {
	userRepo := mocks.NewMockUserRepository()
	tokenService := mocks.NewMockTokenService()
	return userRepo, tokenService, usecases.NewLoadUserByTokenUseCase(userRepo, tokenService)
}

func TestLoadUserByTokenUseCase_Execute(t *testing.T) {
	t.Run("should return user when token is valid", func(t *testing.T) {
		userRepo, tokenService, useCase := setupLoadUserByTokenTest()

		token := "valid-token"
		userID := uuid.New()
		expectedUser, _ := entities.NewUser("John Doe", validEmail(), validPassword(), validPhone(), nil)

		tokenService.On("ValidateToken", token).Return(userID.String(), nil)
		userRepo.On("FindByID", userID.String()).Return(&expectedUser, nil)

		user, err := useCase.Execute(token)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.Name(), user.Name)
		assert.Equal(t, expectedUser.Email().Value(), user.Email)
		assert.Equal(t, expectedUser.Phone().Value(), user.Phone)

		tokenService.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error when token is invalid", func(t *testing.T) {
		_, tokenService, useCase := setupLoadUserByTokenTest()

		token := "invalid-token"
		tokenService.On("ValidateToken", token).Return("", shared.ErrAuthorization)

		user, err := useCase.Execute(token)

		assert.ErrorIs(t, err, shared.ErrAuthorization)
		assert.Nil(t, user)
		tokenService.AssertExpectations(t)
	})

	t.Run("should return error when user is not found", func(t *testing.T) {
		userRepo, tokenService, useCase := setupLoadUserByTokenTest()

		token := "valid-token"
		userID := uuid.New().String()
		tokenService.On("ValidateToken", token).Return(userID, nil)
		userRepo.On("FindByID", userID).Return(nil, shared.ErrNotFound)

		user, err := useCase.Execute(token)

		assert.ErrorIs(t, err, shared.ErrNotFound)
		assert.Nil(t, user)
		tokenService.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})
}
