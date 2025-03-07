package store

import (
	"database/sql"

	"github.com/Dom-HTG/go-template/internal/interfaces"
)

// Store is a struct that holds the database instance.
type Storage struct {
	interfaces.Post
	interfaces.User
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Post: NewPostStore(db),
		User: NewUserStore(db),
	}
}
