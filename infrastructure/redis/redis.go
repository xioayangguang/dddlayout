package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"layout/infrastructure/config"
	"time"
)

var Instances *redis.Client

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Instances.Redis.Addr,
		Password: config.Instances.Redis.Password,
		DB:       config.Instances.Redis.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}
	return rdb
}
