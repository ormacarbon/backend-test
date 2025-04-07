package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/google/uuid"
)

func TestLoadUsersOrderedByPointsUseCase_Execute(t *testing.T) {
	t.Run("successfully returns ranked users", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository()
		useCase := usecases.NewLoadUsersOrderedByPointsUseCase(mockRepo)

		email, _ := object_values.NewEmail("user1@example.com")
		phone, _ := object_values.NewPhoneNumber("+5511999999999")
		password, _ := object_values.NewPassword("StrongP@ssw0rd")
		user1 := entities.LoadUser(uuid.New(), "Alice", email, password, phone, "", nil, 10)

		email2, _ := object_values.NewEmail("user2@example.com")
		phone2, _ := object_values.NewPhoneNumber("+5511988888888")
		password2, _ := object_values.NewPassword("AnotherP@ss")
		user2 := entities.LoadUser(uuid.New(), "Bob", email2, password2, phone2, "", nil, 5)

		mockRepo.On("FindUsersOrderedByPoints", 1, 10).Return([]entities.User{user1, user2}, nil)

		input := dto.GetUsersRankingInput{
			PaginationInput: dto.PaginationInput{
				Page:  1,
				Limit: 10,
			},
		}
		output, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.Len(t, output, 2)
		assert.Equal(t, "Alice", output[0].Name)
		assert.Equal(t, 10, output[0].Points)
		assert.Equal(t, "Bob", output[1].Name)
		assert.Equal(t, 5, output[1].Points)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository returns error", func(t *testing.T) {
		mockRepo := mocks.NewMockUserRepository()
		useCase := usecases.NewLoadUsersOrderedByPointsUseCase(mockRepo)

		mockRepo.On("FindUsersOrderedByPoints", 1, 10).Return(nil, errors.New("db error"))

		input := dto.GetUsersRankingInput{
			PaginationInput: dto.PaginationInput{
				Page:  1,
				Limit: 10,
			},
		}
		output, err := useCase.Execute(input)

		assert.Error(t, err)
		assert.Nil(t, output)
		mockRepo.AssertExpectations(t)
	})
}
