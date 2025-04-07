package factory

import (
	"os"
	"sync"
	"time"

	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	usecases "github.com/cassiusbessa/backend-test/internal/application/use-cases"
	"github.com/cassiusbessa/backend-test/internal/infra/db"
	"github.com/cassiusbessa/backend-test/internal/infra/email"
	"github.com/cassiusbessa/backend-test/internal/infra/token"
	"github.com/cassiusbessa/backend-test/internal/interfaces/http/handlers"
)

type dependencies struct {
	UserRepository output_ports.UserRepository
	TokenService   output_ports.TokenService
	EmailService   output_ports.EmailService
}

var (
	depsInstance *dependencies
	once         sync.Once
)

func getDependencies() *dependencies {
	once.Do(func() {
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "segredo"
		}
		weekDuration := 7 * 24 * time.Hour

		depsInstance = &dependencies{
			UserRepository: db.NewUserGormRepository(db.DB),
			TokenService:   token.NewJWTService(jwtSecret, weekDuration),
			EmailService: email.NewBrevoEmailService(
				"cassiusbessa@gmail.com",
				"bvio",
			),
		}
	})

	return depsInstance
}

func BuildCreateUserHandler() *handlers.CreateUserHandler {
	deps := getDependencies()
	useCase := usecases.NewCreateUserUseCase(deps.UserRepository, deps.EmailService)
	return handlers.NewCreateUserHandler(useCase)
}

func BuildLoginHandler() *handlers.LoginHandler {
	deps := getDependencies()
	useCase := usecases.NewLoginUserUseCase(deps.UserRepository, deps.TokenService)
	return handlers.NewLoginHandler(useCase)
}

func BuildLoadUserByTokenHandler() *handlers.LoadUserByTokenHandler {
	deps := getDependencies()
	useCase := usecases.NewLoadUserByTokenUseCase(deps.UserRepository, deps.TokenService)
	return handlers.NewLoadUserByTokenHandler(useCase)
}

func BuildLoadUsersOrderedByPointsHandler() *handlers.LoadUsersOrderedByPointsHandler {
	deps := getDependencies()
	useCase := usecases.NewLoadUsersOrderedByPointsUseCase(deps.UserRepository)
	return handlers.NewLoadUsersOrderedByPointsHandler(useCase)
}

func BuildFinishCompetitionHandler() *handlers.FinishCompetitionHandler {
	deps := getDependencies()
	useCase := usecases.NewFinishCompetitionUseCase(deps.UserRepository, deps.EmailService)
	return handlers.NewFinishCompetitionHandler(useCase)
}
