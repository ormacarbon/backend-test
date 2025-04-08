package main

import (
	"log"

	"github.com/icl00ud/backend-test/internal/config"
	"github.com/icl00ud/backend-test/internal/routes"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatalf("can't sync zap logger: %v", err)
		}
	}(logger)

	sugar := logger.Sugar()

	sugar.Info("Logger initialized")

	cfg, err := config.LoadConfig()
	if err != nil {
		sugar.Fatalw("Error loading config", "error", err)
	}
	sugar.Info("Configuration loaded")

	app := fiber.New()

	routes.SetupRoutes(app, cfg, sugar)

	listenAddr := ":" + cfg.ServerPort
	sugar.Infow("Starting server", "address", listenAddr)

	if err := app.Listen(listenAddr); err != nil {
		sugar.Fatalw("Failed to start server", "error", err)
	}
}
