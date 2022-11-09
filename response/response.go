package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Response 基类
type Response struct {
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

// ResponseSuccess 成功
func ResponseSuccess(data interface{}) Response {
	if data == nil {
		data = gin.H{}
	}
	return Response{
		Success:   true,
		Timestamp: time.Now().Unix(),
		Code:      SUCCESS,
		Msg:       GetMsg(SUCCESS),
		Data:      data,
	}
}

// ResponseError 错误
func ResponseError(code int) Response {
	return Response{
		Success:   false,
		Timestamp: time.Now().Unix(),
		Code:      code,
		Msg:       GetMsg(code),
		Data:      gin.H{},
	}
}
