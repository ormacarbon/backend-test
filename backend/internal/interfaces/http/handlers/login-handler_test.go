package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoginUserUseCase struct {
	mock.Mock
}

func (m *MockLoginUserUseCase) Execute(input dto.LoginInput) (dto.LoginOutput, error) {
	args := m.Called(input)
	return args.Get(0).(dto.LoginOutput), args.Error(1)
}

func setupLoginHandlerTest() (*MockLoginUserUseCase, *gin.Engine) {
	mockUseCase := new(MockLoginUserUseCase)
	handler := NewLoginHandler(mockUseCase)

	r := gin.Default()
	r.POST("/login", handler.Execute)

	return mockUseCase, r
}

func TestUserHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "john@example.com", Password: "StrongP@ssw0rd"}
		expected := dto.LoginOutput{Token: "fake.jwt.token"}

		mockUseCase.On("Execute", input).Return(expected, nil)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), expected.Token)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		_, router := setupLoginHandlerTest()

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid_json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid request body")
	})

	t.Run("not found error", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "missing@example.com", Password: "123"}
		mockUseCase.On("Execute", input).Return(dto.LoginOutput{}, shared.ErrNotFound)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), shared.ErrNotFound.Error())
	})

	t.Run("unauthorized error", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "user@example.com", Password: "wrong"}
		mockUseCase.On("Execute", input).Return(dto.LoginOutput{}, shared.ErrAuthorization)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), shared.ErrAuthorization.Error())
	})

	t.Run("empty params", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "", Password: ""}
		mockUseCase.On("Execute", input).Return(dto.LoginOutput{}, shared.ErrValidation)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid request body")
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "user@example.com", Password: "123456"}
		mockUseCase.On("Execute", input).Return(dto.LoginOutput{}, shared.ErrInternal)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), shared.ErrInternal.Error())
	})

	t.Run("unexpected error - fallback", func(t *testing.T) {
		mockUseCase, router := setupLoginHandlerTest()

		input := dto.LoginInput{Email: "user@example.com", Password: "123456"}
		mockUseCase.On("Execute", input).Return(dto.LoginOutput{}, errors.New("unknown error"))

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
