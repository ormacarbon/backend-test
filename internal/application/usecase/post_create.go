package usecase

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
)

type PostCreate struct {
	repo database.Repository
}

func NewPostCreate(repo database.Repository) *PostCreate {
	return &PostCreate{repo: repo}
}

func (p *PostCreate) Execute(ctx context.Context, input InputPostCreate, authorID int) error {
	post := &domain.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: authorID,
	}
	err := p.repo.CreatePost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

type InputPostCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
