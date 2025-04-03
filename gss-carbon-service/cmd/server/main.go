package main

import (
	"github.com/icl00ud/backend-test/internal/config"
	"github.com/icl00ud/backend-test/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Loading env vars
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Inicializa o Fiber
	app := fiber.New()

	// Configura as rotas e dependências
	routes.SetupRoutes(app, cfg)

	// Inicia o servidor na porta definida
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
