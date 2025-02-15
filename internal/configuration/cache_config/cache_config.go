package cache_config

import (
	redis "github.com/redis/go-redis/v9"
)

func SetupCacheDB(addr, password string, db int) redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})

	return *client
}