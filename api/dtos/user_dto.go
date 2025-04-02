package dtos

type CreateUserDTO struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	ReferrerCode string `json:"referrer_code"`
}