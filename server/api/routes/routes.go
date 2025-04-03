package routes

import (
	"server/api/handlers"
	"server/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.GET("/leaderboard", userHandler.GetLeaderboard)
	}
}
