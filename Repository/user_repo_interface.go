package repository

import (
	"context"
	entity "golang-clean-architecture/Entity"
)

type UserInterface interface {
	Store(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, id int64) error
	FindId(ctx context.Context, id int64) (entity.User, error)
}
