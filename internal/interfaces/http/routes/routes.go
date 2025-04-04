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
