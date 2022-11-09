package middleware

import (
	"fmt"
	"gin-blog/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 用于处理未知error
		defer func() {
			if err := recover(); err != nil {
				_, file, line, _ := runtime.Caller(3)
				msg := fmt.Sprintf("file=[%s],line=[%d],error=[%s]", file, line, err)
				zap.L().Error(msg, zap.String("middleware", "RecoveryMiddleware"))
				res := response.ResponseError(response.UNKNOWN_ERROR)
				ctx.JSON(http.StatusInternalServerError, res)
				return
			}
		}()
		ctx.Next()
	}
}
