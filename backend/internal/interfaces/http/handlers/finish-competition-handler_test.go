package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFinishCompetitionUseCase struct {
	mock.Mock
}

func (m *MockFinishCompetitionUseCase) Execute() error {
	args := m.Called()
	return args.Error(0)
}

func setupFinishCompetitionTestHandler() (*MockFinishCompetitionUseCase, *gin.Engine) {
	mockUseCase := new(MockFinishCompetitionUseCase)
	handler := NewFinishCompetitionHandler(mockUseCase)

	router := gin.Default()
	router.POST("/admin/competition/finish", handler.Execute)

	return mockUseCase, router
}

func TestFinishCompetitionHandler_Execute(t *testing.T) {
	t.Run("should return 204 when competition is finished successfully", func(t *testing.T) {
		mockUseCase, router := setupFinishCompetitionTestHandler()

		mockUseCase.On("Execute").Return(nil)

		req, _ := http.NewRequest(http.MethodPost, "/admin/competition/finish", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		mockUseCase, router := setupFinishCompetitionTestHandler()

		mockUseCase.On("Execute").Return(assert.AnError)

		req, _ := http.NewRequest(http.MethodPost, "/admin/competition/finish", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockUseCase.AssertExpectations(t)
	})
}
