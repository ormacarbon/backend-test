package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/config"
	"github.com/icl00ud/backend-test/internal/email"
	"github.com/icl00ud/backend-test/internal/handler"
	"github.com/icl00ud/backend-test/internal/middleware"
	"github.com/icl00ud/backend-test/internal/migration"
	"github.com/icl00ud/backend-test/internal/repository"
	"github.com/icl00ud/backend-test/internal/service"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, logger *zap.SugaredLogger) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Fatalw("Failed to connect database", "error", err)
	}

	logger.Info("Database connection established")

	if err := migration.Migrate(db, logger); err != nil {
		logger.Fatalw("Failed to run migrations", "error", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db, logger)

	// --- Services ---
	emailSvc := email.NewEmailService(cfg, logger)
	userService := service.NewUserService(userRepo, emailSvc, logger)
	competitionService := service.NewCompetitionService(userRepo, emailSvc, logger)
	leaderboardService := service.NewLeaderboardService(userRepo, logger)

	// --- Handlers ---
	healthHandler := handler.NewHealthHandler(db, logger)
	userHandler := handler.NewUserHandler(userService, logger)
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardService, logger)
	competitionHandler := handler.NewCompetitionHandler(competitionService, logger)
	referralHandler := handler.NewReferralsHandler(userService, logger)

	app.Use(middleware.SetupCors())

	api := app.Group("/api")

	// Health Check endpoints
	api.Get("/health/ping", healthHandler.Ping)
	api.Get("/health/check", healthHandler.Checker)

	// Endpoints //

	// User
	api.Get("/user/:id", userHandler.GetUserByID)
	api.Get("/user/referral/:token", userHandler.GetUserByReferralToken)
	api.Post("/user/register", userHandler.RegisterUser)

	// Leaderboard
	api.Get("/leaderboard", leaderboardHandler.GetLeaderboard)

	// Competition
	api.Post("/competition/finish", competitionHandler.FinishCompetition)

	// Referrals
	api.Get("/referrals", referralHandler.GetReferrals)
}
