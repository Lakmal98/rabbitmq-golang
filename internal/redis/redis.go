package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis(addr, password string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s", err)
	}
	fmt.Println("Connected to Redis!")
}

func IncrementMessageCount() {
	_, err := redisClient.Incr(ctx, "message_count").Result()
	if err != nil {
		log.Fatalf("Failed to increment message count in Redis: %s", err)
	}
	fmt.Println("Message count incremented in Redis!")
}

func GetMessageCount() int64 {
	count, err := redisClient.Get(ctx, "message_count").Int64()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		log.Fatalf("Failed to get message count from Redis: %s", err)
	}
	return count
}
