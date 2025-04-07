package handlers

import (
	"net/http"

	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	"github.com/gin-gonic/gin"
)

type FinishCompetitionHandler struct {
	useCase input_ports.FinishCompetitionUseCase
}

func NewFinishCompetitionHandler(useCase input_ports.FinishCompetitionUseCase) *FinishCompetitionHandler {
	return &FinishCompetitionHandler{
		useCase: useCase,
	}
}

func (h *FinishCompetitionHandler) Execute(c *gin.Context) {
	if err := h.useCase.Execute(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
