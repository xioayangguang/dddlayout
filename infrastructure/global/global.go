package global

import (
	"github.com/redis/go-redis/v9"
	"layout/infrastructure/config"
)

var (
	GitHash   string
	BuildTime string
	GoVersion string
	Config    *config.Config
	Redis     *redis.Client
)
