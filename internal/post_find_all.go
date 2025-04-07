package internal

import (
	"context"
	"time"
)

type FindPost struct {
	repo Repository
}

func NewFindPost(repo Repository) *FindPost {
	return &FindPost{repo: repo}
}

func (f *FindPost) Execute(ctx context.Context) (*[]OutputFindPost, error) {
	posts, err := f.repo.FindAllPost(ctx)
	if err != nil {
		return nil, err
	}
	var output []OutputFindPost
	for _, post := range posts {
		output = append(output, OutputFindPost{
			ID:           post.ID,
			ReferralCode: post.ReferralCode,
			Title:        post.Title,
			Content:      post.Content,
			AuthorID:     post.AuthorID,
			CreatedAt:    post.CreatedAt,
		})
	}
	return &output, nil
}

type OutputFindPost struct {
	ID           int       `json:"id"`
	ReferralCode string    `json:"referral_code"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	AuthorID     int       `json:"author_id"`
	CreatedAt    time.Time `json:"created_at"`
}
