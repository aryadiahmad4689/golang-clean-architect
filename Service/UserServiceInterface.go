package service

import (
	"context"
	entity "golang-clean-architecture/Entity"
)

type UserServiceInterFace interface {
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
	FindUseer(ctx context.Context, id int64) (entity.User, error)
}
