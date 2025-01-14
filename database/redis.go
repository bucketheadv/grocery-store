package database

import (
	"github.com/go-redis/redis/v8"
	"grocery-store/conf"
)

var RedisClient *redis.Client

func init() {
	config := conf.Config.Redis["main"]
	RedisClient = redis.NewClient(config)
}
