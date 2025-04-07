package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostCreate(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryInMemory()
	mockEventBus := new(MockEventBus)
	createauthor := NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e Event) bool {
		return e.Type == EventTypeIncreasePoint
	})).Return()
	inputAuthor := InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	err := createauthor.Execute(ctx, inputAuthor)
	assert.NoError(t, err)
	postCreate := NewPostCreate(repo)
	inputPost := InputPostCreate{
		Title:   "Post title",
		Content: "Post body",
	}
	postCreate.Execute(ctx, inputPost, 1)
	findPost := NewFindPost(repo)
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
	repo := NewRepositoryInMemory()
	mockEventBus := new(MockEventBus)
	createauthor := NewCreateAuthor(repo, mockEventBus)
	mockEventBus.On("Publish", mock.MatchedBy(func(e Event) bool {
		return e.Type == EventTypeIncreasePoint
	})).Return()
	inputAuthor := InputCreateAuthor{
		Name:  "Jonh Doe",
		Email: "jonhdoe@email.com",
		Phone: "+5511999999999",
	}
	err := createauthor.Execute(ctx, inputAuthor)
	assert.NoError(t, err)
	postCreate := NewPostCreate(repo)
	inputPost := InputPostCreate{
		Title:   "Post title",
		Content: "Post body",
	}
	postCreate.Execute(ctx, inputPost, 1)
	postCreate.Execute(ctx, inputPost, 1)
	postCreate.Execute(ctx, inputPost, 2)
	findPost := NewFindPostByAuthor(repo)
	output, err := findPost.Execute(ctx, 1)
	assert.NoError(t, err)
	posts := *output
	assert.Equal(t, 2, len(posts))
}
