package dto

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Phone    string
}

type CreateUserOutput struct {
	UserID string
}
