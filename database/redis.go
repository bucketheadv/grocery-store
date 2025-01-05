package database

import (
	"HereWeGo/conf"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	config := conf.Config.Redis["main"]
	RedisClient = redis.NewClient(config)
}
