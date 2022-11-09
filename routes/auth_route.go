package routes

import (
	"gin-blog/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(app *gin.Engine) {
	auth := controller.NewUserController()
	app.POST("/api/auth/login", auth.UserLoginController)
}
