package handlers

import (
	"errors"
	"gss-backend/api/presenter"
	"gss-backend/pkg/models"
	services "gss-backend/pkg/services/user"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(user_service services.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody models.User
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		if requestBody.FullName == "" || requestBody.Email == "" || requestBody.PhoneNumber == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(errors.New(
				"Full Name, Email, and Phone Number are required",
			)))
		}

		result, err := user_service.Create(&requestBody)

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
		
	}
}