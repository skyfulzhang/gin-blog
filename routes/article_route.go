package routes

import (
	"gin-blog/controller"
	"gin-blog/middleware"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterArticleRouter(app *gin.Engine) {
	store := persistence.NewInMemoryStore(time.Hour)
	group := app.Group("/api/article", middleware.JWTMiddleware())
	articleController := controller.NewArticleController()
	group.POST("/create", articleController.CreateArticleController)
	group.GET("/list", cache.CachePage(store, 30*time.Minute, articleController.GetArticleListController))
	group.GET("/search", cache.CachePage(store, 30*time.Minute, articleController.FilterArticleByTitleController))
	group.GET("/detail/:id", articleController.GetArticleByIdController)
	group.GET("/category/:id", cache.CachePage(store, 30*time.Minute, articleController.GetCateArticleController))
	group.PUT("/update/:id", articleController.UpdateArticleController)
	group.GET("/:id/on", articleController.PutOnArticleController)
	group.GET("/:id/off", articleController.PullOffArticleController)
	group.GET("/:id/hot", articleController.HotArticleController)
	group.DELETE("/delete/:id", articleController.DeleteArticleController)
}
