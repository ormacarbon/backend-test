package database

import (
	"context"
	"errors"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
)

var ErrAuthorNotFound = errors.New("author not found")
var ErrAuthorExists = errors.New("author already exists")
var ErrPostNotFound = errors.New("post not found")

type Repository interface {
	Save(ctx context.Context, author *domain.Author) error
	Find(ctx context.Context, email string) (*domain.Author, error)
	FindByReferralCode(ctx context.Context, referal string) (*domain.Author, error)
	FindByWinners(ctx context.Context, limit int) ([]domain.Author, error)
	IncreasePoint(ctx context.Context, email string, point uint8) error
	CreatePost(ctx context.Context, post *domain.Post) error
	FindAllPost(ctx context.Context) ([]domain.Post, error)
	FindAllPostByAuthor(ctx context.Context, author int) ([]domain.Post, error)
	FindPostByID(ctx context.Context, authorID int, id int) (*domain.Post, error)
}
