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
		logger.Fatalw("Database connection failed", "error", err)
	}
	logger.Info("Database connected")

	if err := migration.Migrate(db, logger); err != nil {
		logger.Fatalw("Migration failed", "error", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Services
	emailSvc := email.NewEmailService(cfg, logger)
	competitionSvc := service.NewCompetitionService(userRepo, emailSvc, logger)
	leaderboardSvc := service.NewLeaderboardService(userRepo, logger)
	userSvc := service.NewUserService(userRepo, emailSvc, logger)

	// Handlers
	healthHandler := handler.NewHealthHandler(db, logger)
	userHandler := handler.NewUserHandler(userSvc, logger)
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardSvc, logger)
	competitionHandler := handler.NewCompetitionHandler(competitionSvc, logger)
	referralHandler := handler.NewReferralsHandler(userSvc, logger)

	app.Use(middleware.SetupCors())

	api := app.Group("/api")
	api.Get("/health/ping", healthHandler.Ping)
	api.Get("/health/check", healthHandler.Checker)
	api.Get("/user/:id", userHandler.GetUserByID)
	api.Get("/user/referral/:token", userHandler.GetUserByReferralToken)
	api.Post("/user/register", userHandler.RegisterUser)
	api.Get("/leaderboard", leaderboardHandler.GetLeaderboard)
	api.Post("/competition/finish", competitionHandler.FinishCompetition)
	api.Get("/referrals", referralHandler.GetReferrals)
}
