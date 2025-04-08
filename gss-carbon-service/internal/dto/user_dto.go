package dto

import "time"

type RegisterUserRequest struct {
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Phone         string `json:"phone" validate:"required"`
	ReferralToken string `json:"referralToken,omitempty"`
}

// -- Response DTOs ---

type UserResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Points        int       `json:"points"`
	ReferralToken string    `json:"referralToken"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ReducedUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReferralPopulatedResponse struct {
	ID        string              `json:"id"`
	Referrer  ReducedUserResponse `json:"referrer"`
	Referred  ReducedUserResponse `json:"referred"`
	CreatedAt time.Time           `json:"createdAt"`
}
