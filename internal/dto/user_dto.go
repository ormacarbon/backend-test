package dto

type RegisterUserDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	ReferredBy  string `json:"referred_by,omitempty"`
}