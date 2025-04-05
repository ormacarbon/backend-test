package dto

type CreateUserInput struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Password   string  `json:"password" binding:"required"`
	Phone      string  `json:"phone" binding:"required"`
	InviteCode *string `json:"invite_code" binding:"omitempty"`
}

type CreateUserOutput struct {
	UserID string `json:"user_id"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

type LoadedUserOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetUsersRankingInput struct {
	PaginationInput
}

type UserRankingItem struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}
