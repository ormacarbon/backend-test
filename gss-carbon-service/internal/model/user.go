package model

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(100)" json:"name"`
	Email         string    `gorm:"type:varchar(100);unique" json:"email"`
	Phone         string    `gorm:"type:varchar(20)" json:"phone"`
	Points        int       `gorm:"default:1" json:"points"`
	ReferralToken string    `gorm:"type:uuid;unique" json:"referralToken"`
	ReferredBy    *uint     `json:"referredBy,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
