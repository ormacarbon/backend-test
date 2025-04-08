package usecase_test

import (
	"context"
	"testing"

	"github.com/Andreffelipe/carbon_offsets_api/internal/application/usecase"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIncreasePoint(t *testing.T) {
	ctx := context.Background()
	mockEmail := new(usecase.MockEmailService)
	mockEventBus := usecase.NewMockEventBus()
	repo := inmemory.NewRepositoryInMemory()
	createauthor := usecase.NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e eventbus.Event) bool {
		return e.Type == eventbus.EventTypeIncreasePoint
	})).Return()
	input := usecase.InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	err := createauthor.Execute(ctx, input)
	assert.NoError(t, err)
	increasePoint := usecase.NewIncreasePoint(repo, mockEmail)
	inputIncreasePoint := usecase.InputIncreasePoint{
		Referal: "@jonhdoe",
	}
	mockEmail.On("Send", mock.Anything).Return(nil)
	err = increasePoint.Execute(ctx, inputIncreasePoint)
	assert.NoError(t, err)
	findauthor := usecase.NewFindAuthor(repo)
	output, err := findauthor.Execute(ctx, "jonhdoe@email.com")
	assert.NoError(t, err)
	assert.Equal(t, output.Points, uint8(2))
	mockEmail.AssertNumberOfCalls(t, "Send", 1)
}
