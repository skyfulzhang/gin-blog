package middleware

import (
	"bytes"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw ResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw ResponseWriter) WriteString(s string) (int, error) {
	rw.body.WriteString(s)
	return rw.ResponseWriter.WriteString(s)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &ResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyWriter
		// 开始时间
		startTime := time.Now()
		// 请求IP
		clientIP := ctx.ClientIP()
		// 请求方式
		reqMethod := ctx.Request.Method
		// 请求路由
		reqUri := ctx.Request.RequestURI
		// 请求 body
		reqBody := utils.GetRequestData(ctx)
		// 处理请求
		ctx.Next()
		// 结束时间
		endTime := time.Since(startTime)
		// 执行时间
		latencyTime := float64(endTime.Microseconds()) / 1000.0
		// 状态码
		statusCode := ctx.Writer.Status()
		// 获取响应
		responseBody := bodyWriter.body.String()
		// 错误信息
		//errMsg := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		zap.L().Info("LoggerMiddleware",
			zap.String("requestId", ctx.Writer.Header().Get("X-Request-Id")),
			zap.String("clientIP", clientIP),
			zap.String("reqUri", reqUri),
			zap.String("reqMethod", reqMethod),
			zap.String("reqBody", reqBody),
			zap.Int("statusCode", statusCode),
			zap.Float64("latencyTime", latencyTime),
			zap.String("responseBody", responseBody),
			zap.String("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()))
	}
}
