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

type ICategoryController interface {
	CreateCategoryController(c *gin.Context)
	GetCategoryByIdController(c *gin.Context)
	SearchCategoryByNameController(c *gin.Context)
	GetCategoryListController(c *gin.Context)
	UpdateCategoryController(c *gin.Context)
	ActiveCategoryController(c *gin.Context)
	DisableCategoryController(c *gin.Context)
	DeleteCategoryController(c *gin.Context)
}

type CategoryController struct {
	Service service.ICategoryService
}

func NewCategoryController() ICategoryController {
	return CategoryController{Service: service.NewCategoryService()}
}

// GetCategoryByIdController
// @Summary 获取分类详情
// @Description 通过分类id获取分类详情信息
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true "分类id"
// @Success 200 {object} response.Response
// @Router /api/category/detail/{id} [get]
func (cc CategoryController) GetCategoryByIdController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := cc.Service.GetCategoryByIdService(id)
	ctx.JSON(http.StatusOK, res)
}

// GetCategoryListController
// @Summary 获取分类列表
// @Description 默认获取全部分类列表，分页参数为page_num和page_size
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/category/list [get]
func (cc CategoryController) GetCategoryListController(ctx *gin.Context) {
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
	res := cc.Service.GetCategoryListService(pageNum, pageNum)
	ctx.JSON(http.StatusOK, res)
}

type CategoryDto struct {
	Name string `json:"name" binding:"required,min=2,max=64"`
}

// CreateCategoryController
// @Summary 添加分类
// @Description 添加分类，参数为分类name
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body CategoryDto true "分类参数"
// @Success 200 {object} response.Response
// @Router /api/category/create [post]
func (cc CategoryController) CreateCategoryController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var cateIn CategoryDto
	err := ctx.ShouldBindJSON(&cateIn)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "CreateCategoryController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := cc.Service.CreateCategoryService(cateIn.Name)
	ctx.JSON(http.StatusOK, res)
}

// UpdateCategoryController
// @Summary 更新分类
// @Description 更新分类，通过分类id找到分类并更新
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "分类id"
// @Param object body CategoryDto true "分类参数"
// @Success 200 {object} response.Response
// @Router /api/category/update/{id} [put]
func (cc CategoryController) UpdateCategoryController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var cateIn CategoryDto
	id, _ := com.StrTo(ctx.Param("id")).Int()
	err := ctx.ShouldBindJSON(&cateIn)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "UpdateCategoryController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	mapData := utils.StructToMap(cateIn)
	res := cc.Service.UpdateCategoryService(id, mapData)
	ctx.JSON(http.StatusOK, res)
}

// SearchCategoryByNameController
// @Summary 搜索分类
// @Description 通过分类name搜索分类列表
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param title query string  false "分类name"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/category/search [get]
func (cc CategoryController) SearchCategoryByNameController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	name := ctx.Query("name")
	pageNum, _ := com.StrTo(ctx.Query("page")).Int()
	pageSize, _ := com.StrTo(ctx.Query("size")).Int()
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := cc.Service.SearchCategoryService(name, pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)
}

// ActiveCategoryController
// @Summary 激活分类
// @Description 通过分类id激活分类
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "分类id"
// @Success 200 {object} response.Response
// @Router /api/category/enable/{id} [get]
func (cc CategoryController) ActiveCategoryController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := cc.Service.ActiveCategoryService(id)
	ctx.JSON(http.StatusOK, res)
}

// DisableCategoryController
// @Summary 禁用分类
// @Description 通过分类id禁用分类
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "分类id"
// @Success 200 {object} response.Response
// @Router /api/category/disable/{id} [get]
func (cc CategoryController) DisableCategoryController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := cc.Service.DisableCategoryService(id)
	ctx.JSON(http.StatusOK, res)
}

// DeleteCategoryController
// @Summary 删除分类
// @Description 通过分类id删除分类
// @Tags 分类接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "分类id"
// @Success 200 {object} response.Response
// @Router /api/category/delete/{id} [delete]
func (cc CategoryController) DeleteCategoryController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := cc.Service.DeleteCategoryService(id)
	ctx.JSON(http.StatusOK, res)
}
