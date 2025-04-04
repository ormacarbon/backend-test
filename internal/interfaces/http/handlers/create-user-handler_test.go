package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do Use Case
type MockCreateUserUseCase struct {
	mock.Mock
}

func (m *MockCreateUserUseCase) Execute(input dto.CreateUserInput) (dto.CreateUserOutput, error) {
	args := m.Called(input)
	return args.Get(0).(dto.CreateUserOutput), args.Error(1)
}

func setupTestHandler() (*MockCreateUserUseCase, *gin.Engine) {
	mockUseCase := new(MockCreateUserUseCase)
	handler := NewCreateUserHandler(mockUseCase)

	router := gin.Default()
	router.POST("/users", handler.Execute)

	return mockUseCase, router
}

func TestCreateUserHandler_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase, router := setupTestHandler()

		input := dto.CreateUserInput{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "StrongP@ssw0rd",
			Phone:    "+5511987654321",
		}
		output := dto.CreateUserOutput{UserID: "123"}

		mockUseCase.On("Execute", input).Return(output, nil)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		_, router := setupTestHandler()

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("{invalid_json}")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("conflict error", func(t *testing.T) {
		mockUseCase, router := setupTestHandler()

		input := dto.CreateUserInput{
			Name:     "Jane Doe",
			Email:    "jane@example.com",
			Password: "StrongP@ss123",
			Phone:    "+5511981234567",
		}

		mockUseCase.On("Execute", input).Return(dto.CreateUserOutput{}, shared.ErrConflictError)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		mockUseCase.AssertExpectations(t)
	})
}
