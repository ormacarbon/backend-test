package models

import (
	"gorm.io/gorm"
)

// Structs that represents the points the user gains every time someone access the user's referral link
type Points struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	UserId uint `json:"user_id" gorm:"not null; unique"`
	User User `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Points uint `json:"points"`
}