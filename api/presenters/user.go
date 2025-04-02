package presenters

import (
	"gss-backend/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// UserPresenter is a struct that is going to be used to present the user data in a controlled way
// As a minor "business logic", I opted to not return the phone number and the email of the user
type UserPresenter struct {
	ID	   uint   `json:"id"`
	FullName   string `json:"full_name"`
	ReferralCode string `json:"referral_code"`
}

// Functions to transform the user 'raw' data to presenters
func UserSuccessResponse(data *models.User) *fiber.Map {
	user := UserPresenter{
		ID:        data.ID,
		FullName:  data.FullName,
		ReferralCode: data.ReferralCode,

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
			ReferralCode: user.ReferralCode,
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