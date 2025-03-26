package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}