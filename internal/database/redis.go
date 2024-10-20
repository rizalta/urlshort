package database

import "github.com/redis/go-redis/v9"

func InitDB(redisAddr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}
