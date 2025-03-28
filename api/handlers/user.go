package handlers

import (
	"errors"
	"gss-backend/api/dtos"
	"gss-backend/api/presenters"
	services "gss-backend/pkg/services/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(user_service services.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dtos.CreateUserDTO

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		if requestBody.FullName == "" || requestBody.Email == "" || requestBody.PhoneNumber == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New(
				"Full Name, Email, and Phone Number are required",
			)))
		}

		result, err := user_service.Create(&requestBody)

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		return c.JSON(presenters.UserSuccessResponse(result))
		
	}
}

func FindAllUsers(user_service services.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := user_service.FindAll()

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		return c.JSON(presenters.UsersSuccessResponse(result))
	}
}

func FindUserByID(user_service services.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		id, err := strconv.ParseUint(idStr, 10, 32)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		result, err := user_service.FindByID(uint(id))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		return c.JSON(presenters.UserSuccessResponse(result))
	}
}