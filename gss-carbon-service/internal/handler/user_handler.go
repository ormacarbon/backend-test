package handler

import (
	"context"
	"errors"

	"github.com/icl00ud/backend-test/internal/dto"
	"github.com/icl00ud/backend-test/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger.Named("UserHandler"),
	}
}

// RegisterUser register user without refferal token
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	sugar := h.logger.Sugar()
	sugar.Info("Received request to register user")

	var user dto.RegisterUserRequest
	if err := c.BodyParser(&user); err != nil {
		sugar.Warnw("Failed to parse request body for user registration", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	createdUser, err := h.userService.RegisterUser(context.Background(), &user)
	if err != nil {
		sugar.Errorw("User registration service failed", "email", user.Email, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to register user"})
	}

	sugar.Infow("User registration successful", "userID", createdUser.ID, "email", createdUser.Email)
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

// RegisterUserWithReferral register user using a referral token
func (h *UserHandler) RegisterUserWithReferral(c *fiber.Ctx) error {
	sugar := h.logger.Sugar()
	sugar.Info("Received request to register user with referral")

	var user dto.RegisterUserWithReferralRequest
	if err := c.BodyParser(&user); err != nil {
		sugar.Warnw("Failed to parse request body for referral registration", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	createdUser, err := h.userService.RegisterUserWithReferral(context.Background(), &user)
	if err != nil {
		sugar.Errorw("User referral registration service failed", "email", user.Email, "referralToken", user.ReferralToken, "error", err)

		if errors.Is(err, errors.New("invalid referral token")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid referral token"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to register user with referral"})
	}

	sugar.Infow("User referral registration successful", "userID", createdUser.ID, "email", createdUser.Email)
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
