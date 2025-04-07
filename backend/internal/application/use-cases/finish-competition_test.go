package usecases

import (
	"strconv"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/use-cases/mocks"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupFinishCompetitionTest() (*mocks.MockUserRepository, *mocks.MockEmailService, *FinishCompetitionUseCase) {
	repo := mocks.NewMockUserRepository()
	email := mocks.NewMockEmailService()
	usecase := NewFinishCompetitionUseCase(repo, email)
	return repo, email, usecase
}

func TestFinishCompetitionUseCase_Execute(t *testing.T) {
	t.Run("should send emails to top 10 and reset scores", func(t *testing.T) {
		repo, email, usecase := setupFinishCompetitionTest()

		topUsers := []entities.User{}
		for i := range 10 {
			emailStr := "user" + strconv.Itoa(i) + "@example.com"
			phoneStr := "+551199999999" + strconv.Itoa(i)

			u := entities.LoadUser(
				uuid.New(),
				"User",
				mocks.MustEmail(t, emailStr),
				mocks.MustPassword(t, "StrongP@ssw0rd"),
				mocks.MustPhone(t, phoneStr),
				uuid.NewString(),
				nil,
				100-i,
			)
			topUsers = append(topUsers, u)

			email.On(
				"SendEmail",
				u.Email().Value(),
				mock.Anything,
				mock.Anything,
			).Return(nil).Once()
		}

		repo.On("FindUsersOrderedByPoints", 0, 10).Return(topUsers, nil).Once()
		repo.On("ResetAllScores").Return(nil).Once()

		err := usecase.Execute()

		assert.NoError(t, err)
		repo.AssertExpectations(t)
		email.AssertExpectations(t)
	})
}
