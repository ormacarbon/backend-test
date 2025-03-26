package main

import (
	"fmt"
	"gss-backend/api/routes"
	"gss-backend/pkg/config"
	"gss-backend/pkg/models"
	repositories "gss-backend/pkg/repositories/user"
	services "gss-backend/pkg/services/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Loading configuration
	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	// Printing .env variables to see if they are loaded correctly
	fmt.Println(config.POSTGRES_HOST)
	fmt.Println(config.POSTGRES_PORT)
	fmt.Println(config.POSTGRES_USER)
	fmt.Println(config.POSTGRES_PASSWORD)
	fmt.Println(config.POSTGRES_DB)

	// Setting up Postgres DSN
	dsn := fmt.Sprintf((
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC"),
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_DB,
	)

	// Connecting to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	// Setting up Migrations
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error running migrations", err)
	}

	// Setting up Fiber
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GSS Gateway API is up and running! 🚀")
	})

	// Instatiating the User Repo
	userRepo := repositories.NewPostgresUserRepository(db)

	// Setting up the User Service
	userService := services.NewUserService(userRepo)

	api := app.Group("/api")
	routes.UserRouter(api, userService)

	log.Fatal(app.Listen(":3000"))

	
}