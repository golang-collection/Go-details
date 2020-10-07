package middleware

import (
	"Go-details/pkg/app"
	"Go-details/pkg/errcode"
	"Go-details/pkg/limiter"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-09-23 21:00
* @Description: 限流器中间件，防止大量请求压垮服务端
**/

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
