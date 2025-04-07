package internal

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type InputText struct {
	ReferalUsed int
	ReferalLink string
	Author      InputCreateAuthor
}

// Código com WaitGroup para processamento concorrente
func setupAuthorsWithWaitGroup(ctx context.Context, authorsInput []InputText, createauthor *CreateAuthor, increasePoint *IncreasePoint) {
	var wg sync.WaitGroup

	for _, author := range authorsInput {
		wg.Add(1)
		// Captura as variáveis em uma closure para evitar problemas de concorrência
		authorCopy := author
		go func() {
			defer wg.Done()
			// Cria o usuário - esta etapa não é paralelizada internamente
			createauthor.Execute(ctx, authorCopy.Author)
			var pointsWg sync.WaitGroup
			for i := 0; i < authorCopy.ReferalUsed; i++ {
				pointsWg.Add(1)
				go func() {
					defer pointsWg.Done()
					increasePoint.Execute(ctx, InputIncreasePoint{authorCopy.ReferalLink})
				}()
			}
			pointsWg.Wait()
		}()
	}
	wg.Wait()
}

func TestFinalCompetition(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryInMemory()
	mockEmail := new(MockEmailService)
	mockEventBus := NewMockEventBus()
	createauthor := NewCreateAuthor(repo, mockEventBus)
	increasePoint := NewIncreasePoint(repo, mockEmail)

	mockEventBus.On("Publish", mock.MatchedBy(func(e Event) bool {
		return e.Type == EventTypeIncreasePoint
	})).Return()

	authorsInput := []InputText{
		{
			ReferalUsed: 10,
			ReferalLink: "@jonhdoe",
			Author: InputCreateAuthor{
				Name:  "Jonh Doe",
				Email: "jonhdoe@email.com",
				Phone: "+5511999999999",
			},
		},
		{
			ReferalUsed: 8,
			ReferalLink: "@jonhdoe2",
			Author: InputCreateAuthor{
				Name:  "Jonh Doe2",
				Email: "jonhdoe2@email.com",
				Phone: "+5511999999999",
			},
		},
		{
			ReferalUsed: 1,
			ReferalLink: "@jonhdoe3",
			Author: InputCreateAuthor{
				Name:  "Jonh Doe3",
				Email: "jonhdoe3@email.com",
				Phone: "+5511999999999",
			},
		},
		{
			ReferalUsed: 3,
			ReferalLink: "@jonhdoe4",
			Author: InputCreateAuthor{
				Name:  "Jonh Doe4",
				Email: "jonhdoe4@email.com",
				Phone: "+5511999999999",
			},
		},
		{
			ReferalUsed: 6,
			ReferalLink: "@jonhdoe5",
			Author: InputCreateAuthor{
				Name:  "Jonh Doe5",
				Email: "jonhdoe5@email.com",
				Phone: "+5511999999999",
			},
		},
	}

	mockEmail.On("Send", mock.Anything).Return(nil)
	mockEventBus.On("Publish", mock.Anything).Return()
	setupAuthorsWithWaitGroup(ctx, authorsInput, createauthor, increasePoint)

	finalCompetition := NewEndCompetition(repo, mockEmail)
	_, err := finalCompetition.Execute(ctx)
	assert.NoError(t, err)
	mockEventBus.AssertNumberOfCalls(t, "Publish", 5)
	mockEmail.AssertNumberOfCalls(t, "Send", 31)
	repo.Init()
}

func TestEndCompetitionWithoutAuthors(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryInMemory()
	mockEmail := new(MockEmailService)
	mockEmail.On("Send", mock.Anything).Return(nil)
	finalCompetition := NewEndCompetition(repo, mockEmail)
	_, err := finalCompetition.Execute(ctx)
	assert.NoError(t, err)
	mockEmail.AssertNotCalled(t, "Send")
}

func TestEndCompetitionEmailFailure(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryInMemory()
	mockEmail := new(MockEmailService)
	mockEventBus := NewMockEventBus()
	createauthor := NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e Event) bool {
		return e.Type == EventTypeIncreasePoint
	})).Return()
	mockEmail.On("Send", mock.Anything).Return(assert.AnError)
	increasePoint := NewIncreasePoint(repo, mockEmail)
	author := InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	createauthor.Execute(ctx, author)
	increasePoint.Execute(ctx, InputIncreasePoint{"@jonhdoe"})
	finalCompetition := NewEndCompetition(repo, mockEmail)
	_, err := finalCompetition.Execute(ctx)
	assert.Error(t, err)
	mockEmail.AssertCalled(t, "Send", mock.Anything)
}
