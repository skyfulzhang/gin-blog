package middleware

import (
	"gin-blog/response"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"golang.org/x/time/rate"
	"net/http"
)

// 例如 capacity 是100，而 rate 是 0.1，那么每秒会填充10个令牌。
func RateLimitMiddleware() gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(1, 60, 60)
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			res := response.ResponseError(response.ACCESS_LIMIT)
			ctx.JSON(http.StatusForbidden, res)
			return
		}
		ctx.Next()
	}
}

// didip/tollbooth可以根据需要，通过ip、uri、methods、custom headers、basic auth usernames等限流
func LimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 每秒限速 100次,默认使用token令牌
		limiter := tollbooth.NewLimiter(100, nil)
		httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if httpError != nil {
			res := response.ResponseError(response.ACCESS_LIMIT)
			c.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		} else {
			c.Next()
		}
	}
}

// rate限流的使用
func RateLimiter() gin.HandlerFunc {
	//例如： 每秒产生1个令牌，最多存储10个令牌。
	limiter := rate.NewLimiter(1, 10)
	return func(c *gin.Context) {
		//当没有可用的令牌时返回false，也就是当没有可用的令牌时，禁止通行
		if !limiter.Allow() {
			res := response.ResponseError(response.ACCESS_LIMIT)
			c.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}
		//用可用的令牌时放行
		c.Next()
	}
}
