package middleware

import (
	"context"
	"gin-blog/response"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(func(ctx *gin.Context) {
			res := response.ResponseError(response.REQUEST_TIMEOUT)
			ctx.JSON(http.StatusRequestTimeout, res)
		}),
	)
}

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置当前 context 的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		// 覆盖原请求
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
