package routes

import (
	"gin-blog/controller"
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(app *gin.Engine) {
	group := app.Group("/api/user", middleware.JWTMiddleware())
	userController := controller.NewUserController()
	group.POST("/create", userController.CreateUserController)
	group.GET("/list", userController.GetUserListController)
	group.GET("/search", userController.SearchUserByUsernameController)
	group.GET("/detail/:id", userController.GetUserByIdController)
	group.GET("/enable/:id", userController.EnableUserController)
	group.GET("/disable/:id", userController.DisableUserController)
	group.GET("/reset/:id", userController.ResetPwdController)
	group.DELETE("/delete/:id", userController.DeleteUserController)
}
