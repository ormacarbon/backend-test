package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIncreasePoint(t *testing.T) {
	ctx := context.Background()
	mockEmail := new(MockEmailService)
	mockEventBus := NewMockEventBus()
	repo := NewRepositoryInMemory()
	createauthor := NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e Event) bool {
		return e.Type == EventTypeIncreasePoint
	})).Return()
	input := InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	createauthor.Execute(ctx, input)
	increasePoint := NewIncreasePoint(repo, mockEmail)
	inputIncreasePoint := InputIncreasePoint{
		Referal: "@jonhdoe",
	}
	mockEmail.On("Send", mock.Anything).Return(nil)
	increasePoint.Execute(ctx, inputIncreasePoint)
	findauthor := NewFindAuthor(repo)
	output, err := findauthor.Execute(ctx, "jonhdoe@email.com")
	assert.NoError(t, err)
	assert.Equal(t, output.Points, uint8(2))
	mockEmail.AssertNumberOfCalls(t, "Send", 1)
}
