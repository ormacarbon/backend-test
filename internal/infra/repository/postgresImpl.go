package repository

import (
	"context"
	"database/sql"

	"github.com/Andreffelipe/carbon_offsets_api/internal/domain"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/logger"
)

type PostgresImpl struct {
	db  *sql.DB
	log *logger.Logger
}

func NewPostgres(db *sql.DB, log *logger.Logger) *PostgresImpl {
	return &PostgresImpl{
		db:  db,
		log: log,
	}
}

func (p *PostgresImpl) Save(ctx context.Context, author *domain.Author) error {
	p.log.InfoWithFields("Create author", map[string]interface{}{
		"name":          author.Name,
		"email":         author.Email,
		"phone":         author.Phone,
		"points":        author.Points,
		"referral_code": author.ReferralCode,
	})
	query := "INSERT into authors(name,email,phone,points,referral_code) VALUES ($1,$2,$3,$4,$5)"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [Save]", err)
		return err
	}
	_, err = stm.ExecContext(ctx, author.Name, author.Email, author.Phone, author.Points, author.ReferralCode)
	if err != nil {
		p.log.Error("Error to execute statement [Save]", err)
		return err
	}
	return nil
}

func (p *PostgresImpl) Find(ctx context.Context, email string) (*domain.Author, error) {
	p.log.InfoWithFields("find author by email", map[string]interface{}{
		"email": email,
	})
	query := "SELECT * FROM authors WHERE email = $1"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [Find]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, email)
	if err != nil {
		p.log.Error("Error to execute statement [Find]", err)
		return nil, err
	}
	var author domain.Author
	for rows.Next() {
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.Phone, &author.Points, &author.ReferralCode, &author.CreatedAt)
		if err != nil {
			p.log.Error("Error to scan row", err)
			return nil, err
		}
	}
	return &author, nil
}

func (p *PostgresImpl) FindByReferralCode(ctx context.Context, referral_code string) (*domain.Author, error) {
	query := "SELECT email, points FROM authors WHERE referral_code = $1"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [FindByReferralCode]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, referral_code)
	if err != nil {
		p.log.Error("Error to execute statement [FindByReferralCode]", err)
		return nil, err
	}
	var author domain.Author
	for rows.Next() {
		err := rows.Scan(&author.Email, &author.Points)
		if err != nil {
			p.log.Error("Error to scan row [FindByReferralCode]", err)
			return nil, err
		}
	}
	return &author, nil
}

func (p *PostgresImpl) FindByWinners(ctx context.Context, limit int) ([]domain.Author, error) {
	query := "SELECT * FROM authors ORDER BY points DESC LIMIT $1"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [FindByWinners]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, limit)
	if err != nil {
		p.log.Error("Error to execute statement [FindByWinners]", err)
		return nil, err
	}
	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.Phone, &author.Points, &author.ReferralCode, &author.CreatedAt)
		if err != nil {
			p.log.Error("Error to scan row [FindByWinners]", err)
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (p *PostgresImpl) IncreasePoint(ctx context.Context, email string, point uint8) error {
	query := "UPDATE authors SET points = $1 WHERE email = $2"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [IncreasePoint]", err)
		return err
	}
	_, err = stm.ExecContext(ctx, point, email)
	if err != nil {
		p.log.Error("Error to execute statement [IncreasePoint]", err)
		return err
	}
	return nil
}

func (p *PostgresImpl) CreatePost(ctx context.Context, post *domain.Post) error {
	query := "INSERT INTO posts(title,content,author_id) VALUES ($1,$2,$3)"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [CreatePost]", err)
		return err
	}
	_, err = stm.ExecContext(ctx, post.Title, post.Content, post.AuthorID)
	if err != nil {
		p.log.Error("Error to execute statement [CreatePost]", err)
		return err
	}
	return nil
}

func (p *PostgresImpl) FindAllPostByAuthor(ctx context.Context, authorID int) ([]domain.Post, error) {
	query := "SELECT * FROM posts WHERE author_id = $1"
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [FindAllPostByAuthor]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, authorID)
	if err != nil {
		p.log.Error("Error to execute statement [FindAllPostByAuthor]", err)
		return nil, err
	}
	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			p.log.Error("Error to scan row [FindAllPostByAuthor]", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *PostgresImpl) FindPostByID(ctx context.Context, authorID int, id int) (*domain.Post, error) {
	query := `
	SELECT a.referral_code, p.*
	FROM posts p
	INNER JOIN authors a ON p.author_id = a.id
	WHERE p.id = $1 AND p.author_id = $2;
	`
	p.log.InfoWithFields("FindPostByID", map[string]interface{}{
		"id":     id,
		"author": authorID,
		"query":  query,
	})
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [FindPostByID]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx, id, authorID)
	if err != nil {
		p.log.Error("Error to execute statement [FindPostByID]", err)
		return nil, err
	}
	var post domain.Post
	for rows.Next() {
		err := rows.Scan(&post.ReferralCode, &post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			p.log.Error("Error to scan row [FindPostByID]", err)
			return nil, err
		}
	}
	return &post, nil
}

func (p *PostgresImpl) FindAllPost(ctx context.Context) ([]domain.Post, error) {
	query := `
	SELECT a.referral_code, p.*
	FROM posts p
	INNER JOIN authors a ON p.author_id = a.id
	WHERE p.author_id = a.id`
	stm, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		p.log.Error("Error to prepare statement [FindAllPost]", err)
		return nil, err
	}
	rows, err := stm.QueryContext(ctx)
	if err != nil {
		p.log.Error("Error to execute statement [FindAllPost]", err)
		return nil, err
	}
	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ReferralCode, &post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			p.log.Error("Error to scan row [FindAllPost]", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
