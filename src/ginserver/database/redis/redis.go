package redis

import (
	"ginserver/env"

	redis "gopkg.in/redis.v3"
)

func GetRedisClient(db int) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     env.Get("REDIS_HOST"),
		Password: env.Get("REDIS_PASSWORD"),
		DB:       int64(db),
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
