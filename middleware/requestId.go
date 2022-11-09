package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/jtolds/gls"
)

var Mgr = NewContextManager()

func RequestIdMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx.Header("X-Request-Id", requestId)
		ctx.Next()
	}
}
