package service

import (
	"context"
	"errors"
	entity "golang-clean-architecture/Entity"
	repository "golang-clean-architecture/Repository"
)

type userService struct {
	UserRepo repository.UserInterface
}

func NewUserService(repo *repository.UserInterface) UserServiceInterFace {
	return &userService{UserRepo: *repo}
}

func (u *userService) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.UserRepo.Store(ctx, user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userService) DeleteUser(ctx context.Context, id int64) error {

	err := u.UserRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return errors.New("Success Delete")
}

func (u *userService) FindUseer(ctx context.Context, id int64) (entity.User, error) {

	user, err := u.UserRepo.FindId(ctx, id)

	if err != nil {
		return user, err
	}

	return user, nil
}
