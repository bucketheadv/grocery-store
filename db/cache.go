package db

import (
	"context"
	"encoding/json"
	"time"
)

type Function[T any] func() T

func CacheByKey[T any](key string, function Function[T]) T {
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
