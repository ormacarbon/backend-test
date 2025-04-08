package usecase

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
)

type FindPostByAuthor struct {
	repo database.Repository
}

func NewFindPostByAuthor(repo database.Repository) *FindPostByAuthor {
	return &FindPostByAuthor{repo: repo}
}

func (f *FindPostByAuthor) Execute(ctx context.Context, authorID int) (*[]OutputFindPost, error) {
	posts, err := f.repo.FindAllPostByAuthor(ctx, authorID)
	if err != nil {
		return nil, err
	}
	var output []OutputFindPost
	for _, post := range posts {
		output = append(output, OutputFindPost{
			Title:    post.Title,
			Content:  post.Content,
			AuthorID: post.AuthorID,
		})
	}
	return &output, nil
}
