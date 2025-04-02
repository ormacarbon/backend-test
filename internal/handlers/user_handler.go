package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-backend-test/internal/dto"
	"github.com/jpeccia/go-backend-test/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var input dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received referred_by:", input.ReferredBy) // Verifica se o campo foi enviado corretamente

	user, err := u.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "User registered successfully",
		"referral_code": user.ReferralCode,
		"share_link":    fmt.Sprintf("https://test.com/signup?ref=%s", user.ReferralCode),
	})
}
