package handler

import (
	"context"
	"errors"

	"github.com/icl00ud/backend-test/internal/dto"
	"github.com/icl00ud/backend-test/internal/errs"
	"github.com/icl00ud/backend-test/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService service.UserService
	logger      *zap.SugaredLogger
}

func NewUserHandler(userService service.UserService, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger.Named("UserHandler"),
	}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	h.logger.Info("Received request to get user by ID")
	id := c.Params("id")

	user, err := h.userService.GetUserByID(context.Background(), id)
	if err != nil {
		h.logger.Errorw("Failed to get user by ID", "id", id, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get user"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetUserByReferralToken(c *fiber.Ctx) error {
	h.logger.Info("Received request to get user by referral token")
	token := c.Params("token")
	user, err := h.userService.GetUserByReferralToken(context.Background(), token)
	if err != nil {
		h.logger.Errorw("Failed to get user by referral token", "token", token, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get user"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	h.logger.Info("Received request to register user")

	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Warnw("Failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	createdUser, err := h.userService.RegisterUser(context.Background(), &req)
	if err != nil {
		h.logger.Errorw("User registration failed", "email", req.Email, "error", err)
		var appErr *errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.StatusCode).JSON(fiber.Map{"message": appErr.Message})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	h.logger.Info("User registered successfully")
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
