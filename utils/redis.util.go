package utils

import (
	"context"
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/configs"
)

func GetRedis(key string) (*string, error) {
	val, err := configs.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func SetRedis(key string, val interface{}, duration time.Duration) error {
	err := configs.RedisClient.Set(context.Background(), key, val, duration).Err()
	if err != nil {
		return err
	}
	return nil
}
