package routes

import (
	"github.com/icl00ud/backend-test/internal/config"
	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/handler"
	"github.com/icl00ud/backend-test/internal/migration"
	"github.com/icl00ud/backend-test/internal/repository"
	"github.com/icl00ud/backend-test/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Conecta ao Postgres via GORM
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Executa as migrações
	if err := migration.Migrate(db); err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	// Injeta dependências
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	// Instancia o serviço de e-mail (dummy)
	_ = email.NewEmailService()

	// Define os endpoints
	api := app.Group("/api")
	api.Post("/register", userHandler.RegisterUser)
	api.Post("/register/referral", userHandler.RegisterUserWithReferral)
}
