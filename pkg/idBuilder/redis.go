package idBuilder

import (
	"context"
	v9 "github.com/redis/go-redis/v9"
	"layout/infrastructure/redis"
)

func SetRedis(*v9.Client) {
	//todo
}

func Generate(key string, initCallback func() int) int {
	timer := redis.Instances.Incr(context.Background(), key)
	count := timer.Val()
	if count == 1 {
		count = int64(initCallback())
		count++
		redis.Instances.Set(context.Background(), key, count, 0)
	}
	return int(count)
}
