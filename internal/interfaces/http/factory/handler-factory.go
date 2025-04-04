package factory

import (
	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/infra/db"
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/handlers"
)

type UserHandlerFactory struct{}

func NewUserHandlerFactory() *UserHandlerFactory {
	return &UserHandlerFactory{}
}

func (f *UserHandlerFactory) BuildCreateUserHandler() *handlers.CreateUserHandler {
	userRepository := db.NewUserGormRepository(db.DB)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)
	return handlers.NewCreateUserHandler(createUserUseCase)
}
