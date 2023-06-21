package redis

import (
	"context"
	"errors"
	"strconv"

	"github.com/Aldikon/API-user/model"
	"github.com/redis/go-redis/v9"
)

var ErrNotFound = errors.New("not found")

type cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *cache {
	return &cache{
		rdb: rdb,
	}
}

func (c *cache) GetUserByID(ctx context.Context, id uint) (model.User, error) {
	user := model.User{}
	if id == 0 {
		return user, ErrNotFound
	}

	ms := c.rdb.HGetAll(ctx, model.SchemaUser(id))
	err := ms.Scan(&user)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return user, ErrNotFound
		}
		return user, err
	}

	return user, nil
}

func (c *cache) SetUser(ctx context.Context, user model.User) error {
	_, err := c.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		key := model.SchemaUser(user.ID)
		p.HSet(ctx, key, "id", user.ID)
		p.HSet(ctx, key, "name", user.Name)
		p.HSet(ctx, key, "surname", user.Surname)
		p.HSet(ctx, key, "gender", user.Gender)
		p.HSet(ctx, key, "status", user.Status)
		p.HSet(ctx, key, "birth_date", user.BirthDate.String())
		p.HSet(ctx, key, "creat_date", user.CreatDate.String())
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *cache) SetListFilter(ctx context.Context, f model.FilterUser, l []uint) error {
	_, err := c.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		key := f.Schema()
		for _, v := range l {
			p.RPush(ctx, key, v)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *cache) GetListFilter(ctx context.Context, f model.FilterUser) ([]uint, error) {

	key := f.Schema()

	l, err := c.rdb.LLen(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	res := make([]uint, 0, l)

	for i := 0; i < int(l); i++ {
		s, err := c.rdb.LPop(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return nil, ErrNotFound
			}
			return nil, err
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		if n != 0 {
			res = append(res, uint(n))
		}
	}

	return res, nil
}
