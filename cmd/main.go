package main

import (
	"fmt"
	"gss-backend/api/routes"
	"gss-backend/pkg/config"
	"gss-backend/pkg/models"
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"
	userService "gss-backend/pkg/services/user"
	userReferralService "gss-backend/pkg/services/user_referral"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Loading configuration
	log.Println("Loading configuration...")
	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration loaded!")

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
	log.Println("Connecting to the database...")
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database!")

	// Setting up Migrations
	log.Println("Running migrations...")

	err = db.AutoMigrate(&models.User{}, &models.UserReferral{})

	if err != nil {
		log.Fatal("Error running migrations", err)
	}

	log.Println("Migrations complete!")

	// Setting up Fiber
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GSS Gateway API is up and running! 🚀")
	})

	// Instatiating Repositories
	userRepo := userRepo.NewPostgresUserRepository(db)
	userReferralRepo := userReferralRepo.NewPostgresUserReferralRepository(db)

	// Instatiating Services
	userService := userService.NewUserService(userRepo, userReferralRepo)
	userReferralService := userReferralService.NewUserReferralService(userRepo, userReferralRepo)

	// Setting up routes
	api := app.Group("/api")
	routes.UserRouter(api, userService)
	routes.UserReferralRouter(api, userReferralService)

	// Starting the server
	port := fmt.Sprintf(":%s", config.FIBER_PORT) 
	log.Fatal(app.Listen(port))

	
}