package handlers

import (
	"net/http"

	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	"github.com/gin-gonic/gin"
)

type LoadUserByTokenHandler struct {
	uc input_ports.LoadUserByTokenUseCase
}

func NewLoadUserByTokenHandler(uc input_ports.LoadUserByTokenUseCase) *LoadUserByTokenHandler {
	return &LoadUserByTokenHandler{uc: uc}
}

func (h *LoadUserByTokenHandler) Execute(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	user, err := h.uc.Execute(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
