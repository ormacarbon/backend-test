package routes

import (
	"gss-backend/api/handlers"
	services "gss-backend/pkg/services/user"

	"github.com/gofiber/fiber/v2"
)

// UserRouter is a function created to define User related routes
func UserRouter(app fiber.Router, user_service services.IUserService) {
	app.Post("/users", handlers.CreateUser(user_service))
	app.Get("/users", handlers.FindAllUsers(user_service))
	app.Get("/users/:id", handlers.FindUserByID(user_service))
}