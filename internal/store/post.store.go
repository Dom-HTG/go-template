package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"

	"github.com/Dom-HTG/go-template/internal/models"
)

type PostStore struct {
	db *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db}
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, title, created_at
	`

	if e := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		post.Tags,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.Title,
        &post.CreatedAt,
	); e != nil {
		return e
	}

	return nil
}
