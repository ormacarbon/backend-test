package models

import (
	"gorm.io/gorm"
)

type Points struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	UserId uint `json:"user_id"`
	User User `gorm:"foreignKey:UserId"`
	Points uint `json:"points"`
}