package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-backend-test/internal/services"
)

type CompetitionHandler struct {
	competitionService services.CompetitionService
}

func NewCompetitionHandler(service services.CompetitionService) *CompetitionHandler {
	return &CompetitionHandler{competitionService: service}
}

func (h *CompetitionHandler) GetWinners(c *gin.Context) {
	winners, err := h.competitionService.GetTopWinners(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve winners"})
		return
	}

	c.JSON(http.StatusOK, winners)
}

func (h *CompetitionHandler) EndCompetition(c *gin.Context) {
	winners, err := h.competitionService.GetTopWinners(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve winners"})
		return
	}

	err = h.competitionService.NotifyWinners(winners)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send emails to winners"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Competition ended, winners have been notified!"})
}