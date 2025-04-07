package routes

import (
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/factory"
	"github.com/gin-gonic/gin"
)

func WithCreateUser(g *gin.RouterGroup) {
	handler := factory.BuildCreateUserHandler()
	g.POST("/users", handler.Execute)
}

func WithLogin(g *gin.RouterGroup) {
	handler := factory.BuildLoginHandler()
	g.POST("/login", handler.Execute)
}

func WithLoadUserByToken(g *gin.RouterGroup) {
	handler := factory.BuildLoadUserByTokenHandler()
	g.GET("/me", handler.Execute)
}

func WithUsersRanking(g *gin.RouterGroup) {
	handler := factory.BuildLoadUsersOrderedByPointsHandler()
	g.GET("/users/ranking", handler.Execute)
}

func WithFinishCompetition(g *gin.RouterGroup) {
	handler := factory.BuildFinishCompetitionHandler()
	g.POST("/admin/competition/finish", handler.Execute)
}
