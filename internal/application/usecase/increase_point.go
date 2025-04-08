package usecase

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/smtp"
)

type IncreasePoint struct {
	repo database.Repository
	smtp smtp.SendEmail
}

func (i *IncreasePoint) InputIncreasePoint(data eventbus.IncreasePointEventData) InputIncreasePoint {
	return InputIncreasePoint{
		Referal: data.Referal,
	}
}

func NewIncreasePoint(repo database.Repository, smtp smtp.SendEmail) *IncreasePoint {
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
	err = i.smtp.Send(smtp.InputSendEmail{
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
