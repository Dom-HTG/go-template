package interfaces

import (
	"context"

	"github.com/Dom-HTG/go-template/internal/models"
)

type User interface {
	Create(context.Context, *models.User) error
}

type Post interface {
	Create(context.Context, *models.Post) error
}
