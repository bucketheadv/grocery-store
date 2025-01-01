package db

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

func CacheByKey[T any](key string, function func() T) T {
	var redisClient = RedisClient
	var ctx = context.Background()
	value := redisClient.Get(ctx, key)
	if value.Err() == nil {
		var ret T
		err := json.Unmarshal([]byte(value.Val()), &ret)
		if err != nil {
			panic(err)
		}
		return ret
	}
	result := function()
	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	redisClient.Set(ctx, key, bytes, time.Duration(300)*time.Second)
	return result
}

func GetCaches[T any](keys []string) []T {
	var redisClient = RedisClient
	var ctx = context.Background()
	value := redisClient.MGet(ctx, keys...)
	var result = make([]T, 0)
	for _, v := range value.Val() {
		if v == nil {
			continue
		}
		var ret T
		var s = v.(string)
		err := json.Unmarshal(([]byte)(s), &ret)
		if err != nil {
			panic(err)
		}
		result = append(result, ret)
	}
	return result
}

func SetCache(key string, value any, ttl time.Duration) {
	var redisClient = RedisClient
	var ctx = context.Background()
	var s string
	if reflect.TypeOf(value).Kind() == reflect.String {
		s = value.(string)
	} else {
		data, err := json.Marshal(value)
		if err != nil {
			logrus.Error(err)
		}
		s = string(data)
	}
	result := redisClient.Set(ctx, key, s, ttl)
	if result.Err() != nil {
		logrus.Error(result.Err())
	}
}
