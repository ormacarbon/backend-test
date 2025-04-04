package routes

import (
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
)

type UserRouterBuilder struct {
	group             *gin.RouterGroup
	createUserHandler *handlers.CreateUserHandler
}

func NewUserRouterBuilder(group *gin.RouterGroup, createUserHandler *handlers.CreateUserHandler) *UserRouterBuilder {
	return &UserRouterBuilder{
		group:             group,
		createUserHandler: createUserHandler,
	}
}

func (b *UserRouterBuilder) Build() {
	b.group.POST("/", b.createUserHandler.Execute)
}
