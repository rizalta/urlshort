package repo

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	rdb *redis.Client
}

func NewRepo(rdb *redis.Client) *RedisRepo {
	return &RedisRepo{rdb}
}

func (r *RedisRepo) AddURL(ctx context.Context, url, short string) error {
	return r.rdb.Set(ctx, short, url, 0).Err()
}

func (r *RedisRepo) GetURL(ctx context.Context, short string) (string, error) {
	url, err := r.rdb.Get(ctx, short).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
