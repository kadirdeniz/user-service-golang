package tools

import (
	"context"
	"log/slog"
	"user-service-golang/pkg"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	mutex.Lock()
	defer mutex.Unlock()

	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     pkg.Configs.Redis.Host,
			Password: pkg.Configs.Redis.Password,
			DB:       pkg.Configs.Redis.DB,
		})

		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			slog.Error(err.Error())
			panic(err)
		}
	}

	return redisClient
}
