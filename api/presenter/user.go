package presenter

import (
	"gss-backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// UserPresenter is a struct that is going to be used to present the user data in a controlled way
type UserPresenter struct {
	ID	   uint   `json:"id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
}

// Functions to transform the user 'raw' data to presenters
func UserSuccessResponse(data *models.User) *fiber.Map {
	user := UserPresenter{
		ID:        data.ID,
		FullName:  data.FullName,
		Email:     data.Email,
	}

	return &fiber.Map{
		"status": "success",
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]models.User) *fiber.Map {
	var users []UserPresenter

	for _, user := range *data {
		users = append(users, UserPresenter{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
		})
	}

	return &fiber.Map{
		"status": "success",
		"data":   users,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": "error",
		"data":   nil,
		"error":  err.Error(),
	}
}