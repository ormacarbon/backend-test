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

func TestCreateAuthor(t *testing.T) {
	ctx := context.Background()
	repo := inmemory.NewRepositoryInMemory()
	mockEventBus := usecase.NewMockEventBus()
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
	findauthor := usecase.NewFindAuthor(repo)
	output, err := findauthor.Execute(ctx, input.Email)
	assert.NoError(t, err)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)
	assert.Equal(t, input.Phone, output.Phone)
	assert.Equal(t, output.Points, uint8(1))
	assert.Equal(t, output.ReferralCode, "@jonhdoe")
}

// func TestErroAuthorExists(t *testing.T) {
// 	ctx := context.Background()
// 	repo := NewRepositoryInMemory()
// 	mockEventBus := NewMockEventBus()
// 	createauthor := NewCreateAuthor(repo, mockEventBus)
// 	input := InputCreateAuthor{
// 		Name:  "Jonh Doe",
// 		Email: "jonhdoe@email.com",
// 		Phone: "+5511999999999",
// 	}
// 	err := createauthor.Execute(ctx, input)
// 	assert.NoError(t, err)
// 	err = createauthor.Execute(ctx, input)
// 	assert.Equal(t, err, ErrAuthorExists)
// }
