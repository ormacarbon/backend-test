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

func validUserDTO() *dto.LoadedUserOutput {
	return &dto.LoadedUserOutput{
		ID:    "123",
		Name:  "John Doe",
		Email: "valid_email@email.com",
		Phone: "+5511987654321",
	}
}

type LoadUserByTokenUseCaseMock struct {
	mock.Mock
}

func (m *LoadUserByTokenUseCaseMock) Execute(token string) (*dto.LoadedUserOutput, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoadedUserOutput), args.Error(1)
}

func setupLoadUserByTokenHandlerTest() (*LoadUserByTokenUseCaseMock, *gin.Engine) {
	mockUseCase := new(LoadUserByTokenUseCaseMock)
	handler := handlers.NewLoadUserByTokenHandler(mockUseCase)

	router := gin.Default()
	router.GET("/me", handler.Execute)

	return mockUseCase, router
}

func TestLoadUserByTokenHandler_Execute(t *testing.T) {

	t.Run("should return user when token is valid", func(t *testing.T) {
		mockUseCase, router := setupLoadUserByTokenHandlerTest()

		user := validUserDTO()
		mockUseCase.On("Execute", "valid-token").Return(user, nil)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "valid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 401 if token is missing", func(t *testing.T) {
		_, router := setupLoadUserByTokenHandlerTest()

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 if token is invalid", func(t *testing.T) {
		mockUseCase, router := setupLoadUserByTokenHandlerTest()

		mockUseCase.On("Execute", "invalid-token").Return(nil, errors.New("invalid token"))

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "invalid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockUseCase.AssertExpectations(t)
	})
}
