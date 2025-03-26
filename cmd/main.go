package main

import (
	"fmt"
	"gss-backend/pkg/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.POSTGRES_HOST)
	fmt.Println(config.POSTGRES_PORT)
	fmt.Println(config.POSTGRES_USER)
	fmt.Println(config.POSTGRES_PASSWORD)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GSS Gateway API is up and running! 🚀")
	})

	log.Fatal(app.Listen(":3000"))

	
}