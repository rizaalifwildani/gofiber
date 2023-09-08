package configs

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	redisAddress := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,                // Replace with your Redis server address
		Password: os.Getenv("REDIS_PASSWORD"), // No password by default
		DB:       0,                           // Default DB
	})

	// Ping the Redis server to check the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	RedisClient = client
}
