package routes

import (
	_ "gin-blog/docs"
	"gin-blog/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitNoRoute(ctx *gin.Context) {
	res := response.ResponseError(response.PAGE_NOT_FOUND)
	ctx.JSON(http.StatusNotFound, res)
}

func InitNoMethod(ctx *gin.Context) {
	res := response.ResponseError(response.METHOD_NOT_ALLOW)
	ctx.JSON(http.StatusNotFound, res)
}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(InitNoRoute)
	router.NoMethod(InitNoMethod)
	//router.Use(requestid.New())
	RegisterAuthRouter(router)
	RegisterMiddleWare(router)
	RegisterUserRouter(router)
	RegisterArticleRouter(router)
	RegisterCategoryRouter(router)
	RegisterSwagger(router)
	return router
}
