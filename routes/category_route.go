package routes

import (
	"gin-blog/controller"
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRouter(app *gin.Engine) {
	group := app.Group("/api/category", middleware.JWTMiddleware())
	categoryController := controller.NewCategoryController()
	group.POST("/create", categoryController.CreateCategoryController)
	group.GET("/list", categoryController.GetCategoryListController)
	group.GET("/search", categoryController.SearchCategoryByNameController)
	group.GET("/detail/:id", categoryController.GetCategoryByIdController)
	group.PUT("/update/:id", categoryController.UpdateCategoryController)
	group.GET("/enable/:id", categoryController.ActiveCategoryController)
	group.GET("/disable/:id", categoryController.DisableCategoryController)
	group.DELETE("/delete/:id", categoryController.DeleteCategoryController)
}
