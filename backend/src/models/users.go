package models

import (
	"errors"

	"github.com/joaooliveira247/backend-test/src/utils"
)

type Users struct {
	BaseModel
	Name          string `json:"name"  binding:"required,gt=1,lt=256" gorm:"type:varchar(255);not null;column:name"`
	Email         string `json:"email" binding:"required,gt=4,lt=256" gorm:"type:varchar(255);unique;not null;column:email"`
	Phone         string `json:"phone" binding:"required,gt=7,lt=16"  gorm:"type:varchar(16);not null;column:phone"`
	AffiliateCode string `json:"-"                                    gorm:"type:varchar(8);unique;not null;column:affiliate_code"`
}

func (user *Users) Validate() error {
	if !utils.EmailValidator(user.Email) {
		return errors.New("invalid email format")
	} else if !utils.IsPhoneNumber(user.Phone) {
		return errors.New("invalid phone format")
	}
	return nil
}
