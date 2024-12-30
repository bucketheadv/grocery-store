package db

import "github.com/go-redis/redis/v8"

var RedisTemplateClient *redis.Client

func init() {
	RedisTemplateClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
