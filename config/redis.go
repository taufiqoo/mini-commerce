package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect redis", err)
	}

	fmt.Println("Redis Connected", pong)
	return client
}
