package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PhoneNumber  string `json:"phone_number"`
	ReferralCode string `json:"referral_code" gorm:"unique"`
	ReferredBy   string `json:"referred_by"`
	Points       int    `json:"points"`
}
