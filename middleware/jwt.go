package middleware

import (
	"gin-blog/response"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// JWT token验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		token := c.GetHeader("Authorization")
		if token == "" {
			code = response.ERROR_NO_TOKEN
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				zap.L().Error(err.Error(), zap.String("middleware", "JWTMiddleware"))
				code = response.ERROR_INVALID_TOKEN
			} else {
				c.Set("user_id", claims.Id)
				c.Set("user_name", claims.Username)
				code = response.SUCCESS
			}
		}
		if code != response.SUCCESS {
			res := response.ResponseError(code)
			c.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}
		c.Next()
	}
}
