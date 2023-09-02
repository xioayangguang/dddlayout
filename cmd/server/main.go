package main

import (
	"fmt"
	"layout/cmd/server/wireinject"
	"layout/infrastructure/config"
	"layout/infrastructure/http/server"
	"layout/infrastructure/logx"
	"layout/infrastructure/redis"
)

// go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"
var (
	gitHash   string
	buildTime string
	goVersion string
)

// @title YoYo API
// @version 0.0.1
// @description This is a YoYo Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	config.InitConfig()
	redis.Redis = redis.InitRedis()
	engine, cleanup, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	logx.Channel(logx.Default).Info("server start http://127.0.0.1:", config.Config.Http.Port)
	server.Run(engine, fmt.Sprintf(":%d", config.Config.Http.Port))
	defer cleanup()
}
