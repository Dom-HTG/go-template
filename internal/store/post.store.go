package store

import (
	"context"
	"database/sql"
)

type PostStore struct {
	db *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db}
}

func (s *PostStore) Create(ctx context.Context) error {
	return nil
}
