package model

import "context"

type CacheDB interface {
	GetUserByID(ctx context.Context, id uint) (User, error)
	SetUser(ctx context.Context, user User) error
	SetListFilter(ctx context.Context, f FilterUser, l []uint) error
	GetListFilter(ctx context.Context, f FilterUser) ([]uint, error)
}
