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

func TestPostCreate(t *testing.T) {
	ctx := context.Background()
	repo := inmemory.NewRepositoryInMemory()
	mockEventBus := new(usecase.MockEventBus)
	createauthor := usecase.NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e eventbus.Event) bool {
		return e.Type == eventbus.EventTypeIncreasePoint
	})).Return()
	inputAuthor := usecase.InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	err := createauthor.Execute(ctx, inputAuthor)
	assert.NoError(t, err)
	postCreate := usecase.NewPostCreate(repo)
	inputPost := usecase.InputPostCreate{
		Title:   "Post title",
		Content: "Post body",
	}
	err = postCreate.Execute(ctx, inputPost, 1)
	assert.NoError(t, err)
	findPost := usecase.NewFindPost(repo)
	output, err := findPost.Execute(ctx)
	assert.NoError(t, err)
	posts := *output
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, inputPost.Title, posts[0].Title)
	assert.Equal(t, inputPost.Content, posts[0].Content)
	assert.Equal(t, posts[0].AuthorID, 1)
}

func TestPostFindAll(t *testing.T) {
	ctx := context.Background()
	repo := inmemory.NewRepositoryInMemory()
	mockEventBus := new(usecase.MockEventBus)
	createauthor := usecase.NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e eventbus.Event) bool {
		return e.Type == eventbus.EventTypeIncreasePoint
	})).Return()
	inputAuthor := usecase.InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	err := createauthor.Execute(ctx, inputAuthor)
	assert.NoError(t, err)
	postCreate := usecase.NewPostCreate(repo)
	inputPost := usecase.InputPostCreate{
		Title:   "Post title",
		Content: "Post body",
	}
	err = postCreate.Execute(ctx, inputPost, 1)
	assert.NoError(t, err)
	err = postCreate.Execute(ctx, inputPost, 1)
	assert.NoError(t, err)
	findPost := usecase.NewFindPostByAuthor(repo)
	output, err := findPost.Execute(ctx, 1)
	assert.NoError(t, err)
	posts := *output
	assert.Equal(t, 2, len(posts))
}
