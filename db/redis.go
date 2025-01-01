package db

import (
	"HereWeGo/components"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	conf := components.GetConfig().Redis
	RedisClient = redis.NewClient(&conf)
}
