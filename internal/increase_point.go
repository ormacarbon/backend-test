package internal

import (
	"context"
)

type IncreasePoint struct {
	repo Repository
	smtp SendEmail
}

func NewIncreasePoint(repo Repository, smtp SendEmail) *IncreasePoint {
	return &IncreasePoint{
		repo: repo,
		smtp: smtp,
	}
}

func (i *IncreasePoint) Execute(ctx context.Context, input InputIncreasePoint) error {
	author, err := i.repo.FindByReferralCode(ctx, input.Referal)
	if err != nil {
		return err
	}
	points := author.Points + 1
	err = i.repo.IncreasePoint(ctx, author.Email, points)
	if err != nil {
		return err
	}
	err = i.smtp.Send(InputSendEmail{
		To:      author.Email,
		Subject: "Parabéns!",
		Body:    "Você recebeu um ponto em nosso sorteio!",
	})
	if err != nil {
		return err
	}
	return nil
}

type InputIncreasePoint struct {
	Referal string `json:"referal"`
}
