package handlers

import (
	"net/http"
	"strconv"

	"github.com/cassiusbessa/backend-test/internal/application/dto"
	input_ports "github.com/cassiusbessa/backend-test/internal/application/ports/input"
	"github.com/gin-gonic/gin"
)

type LoadUsersOrderedByPointsHandler struct {
	uc input_ports.GetUsersRankingUseCase
}

func NewLoadUsersOrderedByPointsHandler(uc input_ports.GetUsersRankingUseCase) *LoadUsersOrderedByPointsHandler {
	return &LoadUsersOrderedByPointsHandler{uc: uc}
}
func (h *LoadUsersOrderedByPointsHandler) Execute(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if page == "" {
		page = "0"
	}
	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	input := dto.GetUsersRankingInput{
		PaginationInput: dto.PaginationInput{
			Page:  pageInt,
			Limit: limitInt,
		},
	}

	users, err := h.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
