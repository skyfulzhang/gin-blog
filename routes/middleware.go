package routes

import (
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleWare(app *gin.Engine) {
	app.Use(middleware.NoCacheMiddleWare)
	app.Use(middleware.OptionsMiddleWare)
	app.Use(middleware.SecureMiddleWare)
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.RecoveryMiddleware())
	app.Use(middleware.TimeoutMiddleware())
	app.Use(middleware.RequestIdMiddleWare())
	app.Use(middleware.RateLimitMiddleware())
}
