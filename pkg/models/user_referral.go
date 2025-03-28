package models

import (
	"gorm.io/gorm"
)

// Struct that represents the connection
type UserReferral struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	ReferrerId uint `json:"referrer_id" gorm:"not null"` // The user that referred the other user
	Referrer User `gorm:"foreignKey:ReferrerId; constraint:OnDelete:CASCADE"`
	ReferredId uint `json:"referred_id" gorm:"not null"` // The user that was referred
	Referred User `gorm:"foreignKey:ReferredId; constraint:OnDelete:CASCADE"`
}