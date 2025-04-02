package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-backend-test/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, competitionHandler *handlers.CompetitionHandler) {
	api := r.Group("/api")
	{
		api.POST("/signup", userHandler.RegisterUser)
		api.GET("/winners", competitionHandler.GetWinners) // Nova rota para vencedores
	}
}
