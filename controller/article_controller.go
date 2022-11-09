package controller

import (
	"gin-blog/response"
	"gin-blog/service"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

type IArticleController interface {
	CreateArticleController(c *gin.Context)
	GetCateArticleController(c *gin.Context)
	GetArticleByIdController(c *gin.Context)
	FilterArticleByTitleController(c *gin.Context)
	GetArticleListController(c *gin.Context)
	UpdateArticleController(c *gin.Context)
	PutOnArticleController(c *gin.Context)
	PullOffArticleController(c *gin.Context)
	HotArticleController(c *gin.Context)
	DeleteArticleController(c *gin.Context)
}

type ArticleController struct {
	Service service.IArticleService
}

func NewArticleController() ArticleController {
	return ArticleController{Service: service.NewArticleService()}
}

type ArticleIn struct {
	Title   string `json:"title" form:"title" binding:"required,gt=2,lt=64"`
	Author  string `json:"author" form:"author" binding:"required,gt=2,lt=32"`
	Cover   string `json:"cover" form:"cover" binding:"required,gt=2,lt=126"`
	Content string `json:"content" form:"content" binding:"required,gt=2,lt=256"`
	Cid     int    `json:"cid" form:"cid" binding:"required,gt=0"`
}

// CreateArticleController
// @Summary 创建文章
// @Description 创建文章，
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body ArticleIn true "文章参数"
// @Success 200 {object} response.Response
// @Router /api/article/create [post]
func (ac ArticleController) CreateArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var article ArticleIn
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "CreateArticleController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := ac.Service.CreateArticleService(article.Title, article.Author, article.Cover, article.Content, article.Cid)
	ctx.JSON(http.StatusOK, res)
}

// GetArticleByIdController
// @Summary 获取文章详情
// @Description 通过参数文章id获取文章的详情
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "文章id"
// @Success 200 {object} response.Response
// @Router /api/article/detail/{id} [get]
func (ac ArticleController) GetArticleByIdController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := ac.Service.GetArticleByIdService(id)
	ctx.JSON(http.StatusOK, res)
}

// GetCateArticleController
// @Summary 获取分类下的文章列表
// @Description 通过参数分类id获取该id下的文章列表
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "分类id"
// @Success 200 {object} response.Response
// @Router /api/article/category/{id}  [get]
func (ac ArticleController) GetCateArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	pageNum, _ := com.StrTo(ctx.Query("page")).Int()
	pageSize, _ := com.StrTo(ctx.Query("size")).Int()
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := ac.Service.GetCateArticleService(id, pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)
}

// GetArticleListController
// @Summary 获取文章列表
// @Description 默认获取全部文章列表，分页参数为page_num和page_size
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/article/list   [get]
func (ac ArticleController) GetArticleListController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	pageNum, _ := com.StrTo(ctx.Query("page")).Int()
	pageSize, _ := com.StrTo(ctx.Query("size")).Int()
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := ac.Service.GetArticleListService(pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)

}

// UpdateArticleController
// @Summary 更新文章详情
// @Description 通过文章id更新文章详细信息
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "文章id"
// @Param object body ArticleIn true "文章参数"
// @Success 200 {object} response.Response
// @Router /api/article/update/{id}  [put]
func (ac ArticleController) UpdateArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var article ArticleIn
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "UpdateArticleController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	id, _ := com.StrTo(ctx.Param("id")).Int()
	mapData := utils.StructToMap(article)
	res := ac.Service.UpdateArticleService(id, mapData)
	ctx.JSON(http.StatusOK, res)
}

// FilterArticleByTitleController
// @Summary 筛选文章接口
// @Description 条件筛选文章列表，如title、status
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param title query string  false "文章名称"
// @Param content query string  false "文章内容"
// @Param status query integer  false "文章状态"
// @Param is_hot query integer  false "文章热门"
// @Param cid query integer  false "文章分类"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/article/search  [get]
func (ac ArticleController) FilterArticleByTitleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	title := ctx.Query("title")
	content := ctx.Query("content")
	status, _ := com.StrTo(ctx.Query("status")).Int()
	is_hot, _ := com.StrTo(ctx.Query("is_hot")).Int()
	cid, _ := com.StrTo(ctx.Query("cid")).Int()
	pageNum, _ := com.StrTo(ctx.Query("page")).Int()
	pageSize, _ := com.StrTo(ctx.Query("size")).Int()
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := ac.Service.FilterArticleService(title, content, status, is_hot, cid, pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)
}

// PutOnArticleController
// @Summary 上架文章
// @Description 通过文章id将文章上架
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true  "文章id"
// @Success 200 {object} response.Response
// @Router /api/article/{id}/on   [get]
func (ac ArticleController) PutOnArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := ac.Service.PutOnArticleService(id)
	ctx.JSON(http.StatusOK, res)
}

// PullOffArticleController
// @Summary 下架文章
// @Description 通过文章id将文章上架
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true  "文章id"
// @Success 200 {object} response.Response
// @Router /api/article/{id}/off   [get]
func (ac ArticleController) PullOffArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := ac.Service.PullOffArticleService(id)
	ctx.JSON(http.StatusOK, res)
}

// HotArticleController
// @Summary  将文章变为热门
// @Description 通过文章id将文章变为热门
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true  "文章id"
// @Success 200 {object} response.Response
// @Router /api/article/{id}/hot   [get]
func (ac ArticleController) HotArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := ac.Service.HotArticleService(id)
	ctx.JSON(http.StatusOK, res)
}

// DeleteArticleController
// @Summary 删除文章接口
// @Description 通过文章id删除指定的文章
// @Tags 文章接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "文章title"
// @Success 200 {object} response.Response
// @Router /api/article/delete/{id}    [delete]
func (ac ArticleController) DeleteArticleController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := ac.Service.DeleteArticleService(id)
	ctx.JSON(http.StatusOK, res)
}
