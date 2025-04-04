package factory

import (
	"os"
	"time"

	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/infra/db"
	"github.com/cassiusbessa/backend-test/internal/infra/token"
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/handlers"
)

func BuildCreateUserHandler() *handlers.CreateUserHandler {
	userRepository := db.NewUserGormRepository(db.DB)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)
	return handlers.NewCreateUserHandler(createUserUseCase)
}

func BuildLoginHandler() *handlers.LoginHandler {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "segredo"
	}
	weekDuration := time.Hour * 24 * 7
	userRepository := db.NewUserGormRepository(db.DB)
	tokenService := token.NewJWTService(jwtSecret, weekDuration)
	loginUseCase := usecases.NewLoginUserUseCase(userRepository, tokenService)
	return handlers.NewLoginHandler(loginUseCase)
}
