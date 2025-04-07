package internal

import "context"

type PostCreate struct {
	repo Repository
}

func NewPostCreate(repo Repository) *PostCreate {
	return &PostCreate{repo: repo}
}

func (p *PostCreate) Execute(ctx context.Context, input InputPostCreate, authorID int) error {
	post := &Post{
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
