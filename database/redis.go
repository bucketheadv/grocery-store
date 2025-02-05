package database

import (
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/go-redis/redis/v8"
	"grocery-store/conf"
)

var RedisClient *redis.Client

func init() {
	config, ok := conf.Config.Redis["main"]
	if !ok {
		logger.Fatal("未找到 Redis: main 配置")
	}
	RedisClient = redis.NewClient(&config)
}
