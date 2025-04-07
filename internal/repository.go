package internal

import (
	"context"
	"errors"
)

var ErrAuthorNotFound = errors.New("author not found")
var ErrAuthorExists = errors.New("author already exists")
var ErrPostNotFound = errors.New("post not found")

type Repository interface {
	Save(ctx context.Context, author *Author) error
	Find(ctx context.Context, email string) (*Author, error)
	FindByReferralCode(ctx context.Context, referal string) (*Author, error)
	FindByWinners(ctx context.Context, limit int) ([]Author, error)
	IncreasePoint(ctx context.Context, email string, point uint8) error
	CreatePost(ctx context.Context, post *Post) error
	FindAllPost(ctx context.Context) ([]Post, error)
	FindAllPostByAuthor(ctx context.Context, author int) ([]Post, error)
	FindPostByID(ctx context.Context, authorID int, id int) (*Post, error)
}
