package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type GetUsersRankingUseCaseMock struct {
	mock.Mock
}

func (m *GetUsersRankingUseCaseMock) Execute(input dto.GetUsersRankingInput) ([]dto.UserRankingItem, error) {
	args := m.Called(input)
	return args.Get(0).([]dto.UserRankingItem), args.Error(1)
}

func validUserRankingList() []dto.UserRankingItem {
	return []dto.UserRankingItem{
		{UserID: "1", Name: "Alice", Points: 100},
		{UserID: "2", Name: "Bob", Points: 80},
	}
}

func setupLoadUsersOrderedByPointsHandlerTest() (*GetUsersRankingUseCaseMock, *gin.Engine) {
	mockUseCase := new(GetUsersRankingUseCaseMock)
	handler := handlers.NewLoadUsersOrderedByPointsHandler(mockUseCase)

	router := gin.Default()
	router.GET("/ranking", handler.Execute)

	return mockUseCase, router
}

func TestLoadUsersOrderedByPointsHandler_Execute(t *testing.T) {

	t.Run("should return 200 with user ranking list", func(t *testing.T) {
		mockUseCase, router := setupLoadUsersOrderedByPointsHandlerTest()

		expectedInput := dto.GetUsersRankingInput{
			PaginationInput: dto.PaginationInput{
				Page:  1,
				Limit: 10,
			},
		}
		expectedOutput := validUserRankingList()

		mockUseCase.On("Execute", expectedInput).Return(expectedOutput, nil)

		req := httptest.NewRequest(http.MethodGet, "/ranking?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 400 if page is invalid", func(t *testing.T) {
		_, router := setupLoadUsersOrderedByPointsHandlerTest()

		req := httptest.NewRequest(http.MethodGet, "/ranking?page=abc&limit=10", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 if limit is invalid", func(t *testing.T) {
		_, router := setupLoadUsersOrderedByPointsHandlerTest()

		req := httptest.NewRequest(http.MethodGet, "/ranking?page=1&limit=xyz", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 if usecase fails", func(t *testing.T) {
		mockUseCase, router := setupLoadUsersOrderedByPointsHandlerTest()

		expectedInput := dto.GetUsersRankingInput{
			PaginationInput: dto.PaginationInput{
				Page:  1,
				Limit: 10,
			},
		}

		mockUseCase.On("Execute", expectedInput).Return(nil, errors.New("internal error"))

		req := httptest.NewRequest(http.MethodGet, "/ranking?page=1&limit=10", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockUseCase.AssertExpectations(t)
	})
}
