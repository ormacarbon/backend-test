package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type InputText struct {
	ReferalUsed int
	ReferalLink string
	Author      InputCreateAuthor
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

	for _, author := range authorsInput {
		createauthor.Execute(ctx, author.Author)
		for i := 0; i < author.ReferalUsed; i++ {
			increasePoint.Execute(ctx, InputIncreasePoint{author.ReferalLink})
		}
	}
	finalCompetition := NewEndCompetition(repo, mockEmail)
	_, err := finalCompetition.Execute(ctx)
	assert.NoError(t, err)
	mockEventBus.AssertNumberOfCalls(t, "Publish", 5)
	mockEmail.AssertNumberOfCalls(t, "Send", 33)
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
