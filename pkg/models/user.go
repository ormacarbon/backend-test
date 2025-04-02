package models

import (
	"gorm.io/gorm"
)

// User struct to model the user entity
type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	FullName string `json:"full_name" gor:"not null"`
	Email string `json:"email" gorm:"not null; unique"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	ReferralCode string `json:"referral_code" gorm:"not null; unique"`
}