package db

import (
	"HereWeGo/initializers"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	conf := initializers.GetConfig().Redis
	RedisClient = redis.NewClient(&conf)
}
