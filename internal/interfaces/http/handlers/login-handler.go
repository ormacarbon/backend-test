package handlers

import (
	"net/http"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginUserUseCase input_ports.LoginUseCase
}

func NewLoginHandler(loginUserCase input_ports.LoginUseCase) *LoginHandler {
	return &LoginHandler{
		loginUserUseCase: loginUserCase,
	}
}

func (h *LoginHandler) Execute(ctx *gin.Context) {
	var input dto.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	output, err := h.loginUserUseCase.Execute(input)
	if err != nil {
		switch err {
		case shared.ErrNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		case shared.ErrAuthorization:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		case shared.ErrValidation:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		case shared.ErrConflictError:
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		case shared.ErrInternal:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		default:
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": output.Token,
	})
}
