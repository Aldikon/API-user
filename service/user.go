package service

import (
	"context"
	"fmt"

	"github.com/Aldikon/API-user/model"
)

type user struct {
	db model.UserRepo
}

func NewUser(db model.UserRepo) *user {
	return &user{
		db: db,
	}
}

func (u *user) Creat(ctx context.Context, user model.User) error {
	err := user.Validation()
	if err != nil {
		return err
	}

	err = u.db.Creat(ctx, user)
	if err != nil {
		return fmt.Errorf("repo Create err: %w", err)
	}

	return nil
}

func (u *user) GetByID(ctx context.Context, id uint) (model.User, error) {
	user := model.User{}
	var err error

	if id == 0 {
		return user, model.NewValidationError("not correctly id")
	}

	user, err = u.db.GetByID(ctx, id)
	if err != nil {
		return user, fmt.Errorf("repo GetById err: %w", err)
	}

	return user, nil
}

func (u *user) GetByFilter(ctx context.Context, f model.FilterUser) ([]model.User, error) {
	err := f.Validate()
	if err != nil {
		return nil, err
	}

	users, err := u.db.GetByFilter(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("repo GetByFilter err: %w", err)
	}

	return users, nil
}

func (u *user) Update(ctx context.Context, user model.User) error {
	err := user.Validation()
	if err != nil {
		return err
	}

	err = u.db.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("repo Update err: %w", err)
	}

	return nil
}

func (u *user) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return model.NewValidationError("not correctly id")
	}

	err := u.db.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("repo Delete err: %w", err)
	}
	return nil
}
