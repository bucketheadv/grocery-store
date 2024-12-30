package db

import (
	"context"
	"encoding/json"
	"time"
)

type Function[T any] func() T

func CacheByKey[T any](key string, function Function[T]) T {
	var redisTemplate = RedisTemplateClient
	var ctx = context.Background()
	value := redisTemplate.Get(ctx, key)
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
	redisTemplate.Set(ctx, key, bytes, time.Duration(300)*time.Second)
	return result
}
