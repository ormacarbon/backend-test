package handlers

import (
	"net/http"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase input_ports.CreateUserUseCase
}

func NewUserHandler(createUserUseCase input_ports.CreateUserUseCase) *UserHandler {
	return &UserHandler{createUserUseCase: createUserUseCase}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var input dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	output, err := h.createUserUseCase.Execute(input)
	switch err {
	case shared.ErrNotFound:
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case shared.ErrAuthorization:
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case shared.ErrValidation:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case shared.ErrConflictError:
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case shared.ErrInternal:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	default:
	}

	ctx.JSON(http.StatusCreated, output)
}
