package model

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(100)" json:"name"`
	Email         string    `gorm:"type:varchar(100);unique" json:"email"`
	Phone         string    `gorm:"type:varchar(20)" json:"phone"`
	Points        int       `gorm:"default:1" json:"points"`
	ReferralToken string    `gorm:"type:uuid;unique" json:"referral_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
