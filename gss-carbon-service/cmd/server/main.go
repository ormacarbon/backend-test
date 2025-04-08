package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/icl00ud/backend-test/internal/config"
	"github.com/icl00ud/backend-test/internal/routes"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatalf("cannot sync logger: %v", err)
		}
	}(logger)
	sugar := logger.Sugar()
	sugar.Info("Logger initialized")

	cfg, err := config.LoadConfig()
	if err != nil {
		sugar.Fatalw("Error loading configuration", "error", err)
	}
	sugar.Info("Configuration loaded")

	app := fiber.New()

	routes.SetupRoutes(app, cfg, sugar)

	go func() {
		listenAddr := ":" + cfg.ServerPort
		sugar.Infow("Starting server", "address", listenAddr)
		if err := app.Listen(listenAddr); err != nil {
			sugar.Fatalw("Failed to start server", "error", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	sugar.Info("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		sugar.Errorw("Server forced to shutdown", "error", err)
	}

	sugar.Info("Server exiting")
}
