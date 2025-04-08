package controllers

import (
	"fmt"
	"server/internal/models"
	"server/internal/repository"
	"server/utils"

	"github.com/google/uuid"
)

type UserController struct {
	userRepo repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func (c *UserController) RegisterUser(name, email, phone, referralCode string) (*models.User, error) {
	user := &models.User{
		Name:   name,
		Email:  email,
		Phone:  phone,
		Points: 1,
	}

	if referralCode != "" {
		referrer, err := c.userRepo.FindByShareCode(referralCode)
		if err == nil {
			user.ReferredBy = referrer.ShareCode
			c.userRepo.UpdatePoints(referrer, referrer.Points+1)

			go utils.SendEmail(referrer.Email, "You earned a point!",
				fmt.Sprintf("You earned a point! %s (%s) signed up using your code.", user.Name, user.Email))
		}
	}

	user.ShareCode = uuid.New().String()

	if err := c.userRepo.Create(user); err != nil {
		return nil, err
	}

	go utils.SendEmail(user.Email, "Welcome to Carbon Offsets Awareness Program!",
		"Thank you for joining us! Your share code is: "+user.ShareCode)

	return user, nil
}

func (c *UserController) GetLeaderboard(sort string, search string, page int) ([]models.User, int64, error) {
	filters := repository.Filters{
		Sort:   sort,
		Search: search,
		Page:   page,
		Limit:  10,
	}
	return c.userRepo.GetLeaderboard(filters)
}

func (c *UserController) GetUserByID(id uint) (*models.User, error) {
	user, err := c.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
