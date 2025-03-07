package interfaces

import "context"

type User interface {
	Create(context.Context) error
}

type Post interface {
	Create(context.Context) error
}
