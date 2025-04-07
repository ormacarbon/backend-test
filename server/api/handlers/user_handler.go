package handlers

import (
	"math"
	"net/http"
	"server/internal/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userController *controllers.UserController
}

func NewUserHandler(userController *controllers.UserController) *UserHandler {
	return &UserHandler{userController: userController}
}

type RegisterRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone_number" binding:"required"`
	ReferralCode string `json:"referral_code,omitempty"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userController.RegisterUser(req.Name, req.Email, req.Phone, req.ReferralCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	sort := c.DefaultQuery("sort", "points")
	search := c.DefaultQuery("search", "")
	page := 1
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	users, total, err := h.userController.GetLeaderboard(sort, search, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve leaderboard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": users,
		"total":       total,
		"page":        page,
		"totalPages":  int(math.Ceil(float64(total) / 10.0)),
	})
}
