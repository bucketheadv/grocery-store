package db

import (
	"HereWeGo/initializers"
	"github.com/go-redis/redis/v8"
)

var RedisTemplateClient *redis.Client

func init() {
	conf := initializers.GetConfig().Redis
	RedisTemplateClient = redis.NewClient(&conf)
}
