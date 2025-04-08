package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email" gorm:"unique"`
	Phone      string `json:"phone" binding:"required"`
	Points     int    `json:"points" gorm:"default:0"`
	ShareCode  string `json:"share_code" gorm:"unique"`
	ReferredBy string `json:"referred_by,omitempty"`
}
