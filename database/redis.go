package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"grocery-store/conf"
)

var RedisClient *redis.Client

func init() {
	config, ok := conf.Config.Redis["main"]
	if !ok {
		logrus.Fatalln("未找到 Redis: main 配置")
	}
	RedisClient = redis.NewClient(&config)
}
