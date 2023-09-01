package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	response2 "layout/infrastructure/response"
	"layout/infrastructure/speedLimit"
)

// SpeedLimit 分布式限速，要是网关做了ip哈希可以直接本地限速提高效率，减少网络io
func SpeedLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiAuth := c.Request.Header.Get("ApiAuth")
		if ApiAuth != "" && speedLimit.SpeedLimit(c, ApiAuth, 1, 10) {
			response2.FailWithCode(c, response2.RateIsTooHigh)
			c.Abort()
			return
		}
		ip := c.ClientIP()
		if speedLimit.SpeedLimit(c, ip, 1, 10) {
			response2.FailWithCode(c, response2.RateIsTooHigh)
			c.Abort()
			return
		}
		c.Next()
	}
}

var limiters = speedLimit.NewLimiters()

// TokenLimit 单机限速
func TokenLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.RemoteIP()
		l := limiters.GetLimiter(ip, rate.Limit(10), 2)
		if !l.Allow() {
			response2.FailWithCode(ctx, response2.RateIsTooHigh)
			ctx.Abort()
		}
		ctx.Next()
	}
}
