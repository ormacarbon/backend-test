package model

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name"`
	Email         string    `gorm:"uniqueIndex" json:"email"`
	Phone         string    `json:"phone"`
	Points        int       `json:"points"`
	ReferralToken string    `gorm:"uniqueIndex" json:"referralToken"`
	ReferredBy    *uint     `json:"referredBy,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
