package internal

import (
	"context"
	"fmt"
)

const MaxWinnersLenth = 10

type EndCompetition struct {
	repo Repository
	smtp SendEmail
}

func NewEndCompetition(repo Repository, smtp SendEmail) *EndCompetition {
	return &EndCompetition{repo: repo, smtp: smtp}
}

func (f *EndCompetition) Execute(ctx context.Context) (*EndCompetitionOutput, error) {
	authors, err := f.repo.FindByWinners(ctx, MaxWinnersLenth)
	if err != nil {
		return nil, err
	}
	// var wg sync.WaitGroup
	// wg.Add(len(authors))
	// for position, author := range authors {
	// 	go func(position int, author Author) {
	// 		defer wg.Done()
	// 		err := f.smtp.Send(InputSendEmail{
	// 			To:      author.Email,
	// 			Subject: "Parabéns! Você está entre os 3 primeiros!\n",
	// 			Body: fmt.Sprintf("Olá %s \n", author.Name) +
	// 				"Parabéns por sua conquista na competição Compensações de Carbono! \n\n" +

	// 				fmt.Sprintf("Sua posição na competição: %d \n", position+1) +
	// 				fmt.Sprintf("Sua pontuação: %d \n\n", author.Points) +

	// 				"Continue participando de nossas competições! \n\n" +

	// 				"Atenciosamente, \n" +
	// 				"Equipe de Competições \n",
	// 		})
	// 		if err != nil {
	// 			return
	// 		}
	// 	}(position, author)
	// }
	// wg.Wait()
	for position, author := range authors {
		err := f.smtp.Send(InputSendEmail{
			To:      author.Email,
			Subject: "Parabéns! Você está entre os 3 primeiros!\n",
			Body: fmt.Sprintf("Olá %s \n", author.Name) +
				"Parabéns por sua conquista na competição Compensações de Carbono! \n\n" +

				fmt.Sprintf("Sua posição na competição: %d \n", position+1) +
				fmt.Sprintf("Sua pontuação: %d \n\n", author.Points) +

				"Continue participando de nossas competições! \n\n" +

				"Atenciosamente, \n" +
				"Equipe de Competições \n",
		})
		if err != nil {
			return nil, err
		}
	}
	return &EndCompetitionOutput{Winners: authors}, nil
}

type EndCompetitionOutput struct {
	Winners []Author `json:"winners"`
}
