package inmemory

import (
	"context"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
)

type RepositoryInMemory struct {
	authors []domain.Author
	posts   []domain.Post
}

var instance *RepositoryInMemory

func NewRepositoryInMemory() *RepositoryInMemory {
	instance = new(RepositoryInMemory)
	instance.Init()
	return instance
}

func (r *RepositoryInMemory) Init() {
	r.authors = []domain.Author{}
	r.posts = []domain.Post{}
}

func (r *RepositoryInMemory) Save(ctx context.Context, author *domain.Author) error {
	author.ID = len(r.authors) + 1
	r.authors = append(r.authors, *author)
	return nil
}

func (r *RepositoryInMemory) Find(ctx context.Context, email string) (*domain.Author, error) {
	for _, author := range r.authors {
		if author.Email == email {
			return &author, nil
		}
	}
	return &domain.Author{}, database.ErrAuthorNotFound
}

func (r *RepositoryInMemory) FindByReferralCode(ctx context.Context, referal string) (*domain.Author, error) {
	for _, author := range r.authors {
		if author.ReferralCode == referal {
			return &author, nil
		}
	}
	return nil, database.ErrAuthorNotFound
}

func (r *RepositoryInMemory) FindByWinners(ctx context.Context, limit int) ([]domain.Author, error) {
	authorsCopy := make([]domain.Author, len(r.authors))
	copy(authorsCopy, r.authors)
	for i := 0; i < len(authorsCopy); i++ {
		for j := i + 1; j < len(authorsCopy); j++ {
			if authorsCopy[i].Points < authorsCopy[j].Points {
				authorsCopy[i], authorsCopy[j] = authorsCopy[j], authorsCopy[i]
			}
		}
	}
	if len(authorsCopy) > limit {
		authorsCopy = authorsCopy[:limit]
	}
	return authorsCopy, nil
}

func (r *RepositoryInMemory) IncreasePoint(ctx context.Context, email string, point uint8) error {
	for i, author := range r.authors {
		if author.Email == email {
			r.authors[i].Points = point
			return nil
		}
	}
	return database.ErrAuthorNotFound
}

func (r *RepositoryInMemory) CreatePost(ctx context.Context, post *domain.Post) error {
	for _, author := range r.authors {
		if author.ID == post.AuthorID {
			r.posts = append(r.posts, *post)
			return nil
		}
	}
	return database.ErrAuthorNotFound
}

func (r *RepositoryInMemory) FindAllPost(ctx context.Context) ([]domain.Post, error) {
	return r.posts, nil
}

func (r *RepositoryInMemory) FindAllPostByAuthor(ctx context.Context, author int) ([]domain.Post, error) {
	var posts []domain.Post
	for _, post := range r.posts {
		if post.AuthorID == author {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func (r *RepositoryInMemory) FindPostByID(ctx context.Context, authorID int, id int) (*domain.Post, error) {
	for _, post := range r.posts {
		if post.ID == id && post.AuthorID == authorID {
			return &post, nil
		}
	}
	return nil, database.ErrPostNotFound
}
