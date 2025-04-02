package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-backend-test/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	r.POST("/signup", userHandler.RegisterUser)
}