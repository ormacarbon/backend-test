package internal

import "context"

type FindAuthor struct {
	repo Repository
}

func NewFindAuthor(repo Repository) *FindAuthor {
	return &FindAuthor{repo: repo}
}

func (f *FindAuthor) Execute(ctx context.Context, email string) (*OutputFindAuthor, error) {
	author, err := f.repo.Find(ctx, email)
	if err != nil {
		return nil, err
	}
	return &OutputFindAuthor{
		Name:         author.Name,
		Email:        author.Email,
		Phone:        author.Phone,
		Points:       author.Points,
		ReferralCode: author.ReferralCode,
	}, nil
}

type OutputFindAuthor struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Points       uint8  `json:"points"`
	ReferralCode string `json:"referral_code"`
}
