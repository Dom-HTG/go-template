package store

import (
	"context"
	"database/sql"

	"github.com/Dom-HTG/go-template/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db}
}

func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING  id, email
	`
	if e := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.Email,
	); e != nil {
		return e
	}

	return nil
}
