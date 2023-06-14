package db

import (
	"context"

	"github.com/Aldikon/API-user/model"
	"github.com/jackc/pgx/v5"
)

type userDB struct {
	db *pgx.Conn
}

func NewUser(db *pgx.Conn) *userDB {
	return &userDB{
		db: db,
	}
}

func (u *userDB) Creat(ctx context.Context, user model.User) error {
	return nil
}

func (u *userDB) GetByID(ctx context.Context, id uint) (model.User, error) {
	user := model.User{}
	return user, nil
}

func (u *userDB) GetByFilter(ctx context.Context, f model.FilterUser) ([]model.User, error) {
	return nil, nil
}

func (u *userDB) Update(ctx context.Context, user model.User) error {
	return nil
}

func (u *userDB) Delete(ctx context.Context, id uint) error {
	return nil
}
