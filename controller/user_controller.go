package controller

import (
	"fmt"
	"gin-blog/response"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

type IUserController interface {
	UserLoginController(c *gin.Context)
	CreateUserController(c *gin.Context)
	GetUserByIdController(c *gin.Context)
	GetUserListController(c *gin.Context)
	SearchUserByUsernameController(c *gin.Context)
	EnableUserController(c *gin.Context)
	DisableUserController(c *gin.Context)
	ResetPwdController(c *gin.Context)
	DeleteUserController(c *gin.Context)
}

type UserController struct {
	Service service.IUserService
}

func NewUserController() IUserController {
	return UserController{Service: service.NewUserService()}
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDto struct {
	Username string `json:"username" form:"username" binding:"required,gt=2,lt=64"`
	Password string `json:"password" form:"password" binding:"required,gt=2,lt=64"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required,gt=2,lt=64"`
	Phone    string `json:"phone" form:"phone" binding:"required,gt=2,lt=64"`
	Email    string `json:"email" form:"email" binding:"required,gt=2,lt=64"`
}

// UserLoginController
// @Summary 用户登录
// @Description 用户登录，返回token
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param object body LoginDto true "登录参数"
// @Success 200 {object} response.Response
// @Router /api/auth/login [post]
func (uc UserController) UserLoginController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var login LoginDto
	err := ctx.ShouldBind(&login)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "UserLoginController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := uc.Service.UserLoginService(login.Username, login.Password)
	ctx.JSON(http.StatusOK, res)
}

// CreateUserController
// @Summary 添加用户
// @Description 添加用户，参数为username和password
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body UserDto true "用户参数"
// @Success 200 {object} response.Response
// @Router /api/user/create [post]
func (uc UserController) CreateUserController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	var user UserDto
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		zap.L().Error(ctx.Request.RequestURI, zap.String("controller", "CreateUserController"), zap.Any("error", err))
		res := response.ResponseError(response.INVALID_PARAMS)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := uc.Service.CreateUserService(user.Username, user.Password, user.Avatar, user.Phone, user.Email)
	ctx.JSON(http.StatusOK, res)

}

// GetUserByIdController
// @Summary 获取用户详情页
// @Description 通过用户id获取用户详情
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true "用户id"
// @Success 200 {object} response.Response
// @Router /api/user/detail/{id} [get]
func (uc UserController) GetUserByIdController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := uc.Service.GetUserByIdService(id)
	ctx.JSON(http.StatusOK, res)

}

// GetUserListController
// @Summary 获取用户列表
// @Description 默认获取全部用户列表，分页参数为page_num和page_size
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/user/list [get]
func (uc UserController) GetUserListController(ctx *gin.Context) {
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
	res := uc.Service.GetUserListService(pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)
}

// SearchUserByUsernameController
// @Summary 搜索用户列表
// @Description 通过用户username获取用户列表数据
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username query string  false "用户username"
// @Param page_num query integer  false "当前第X页"
// @Param page_size query integer  false "每页数量"
// @Success 200 {object} response.Response
// @Router /api/user/list [get]
func (uc UserController) SearchUserByUsernameController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	username := ctx.Query("username")
	pageNum, _ := com.StrTo(ctx.Query("page")).Int()
	pageSize, _ := com.StrTo(ctx.Query("size")).Int()
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := uc.Service.SearchUserService(username, pageNum, pageSize)
	ctx.JSON(http.StatusOK, res)
}

// EnableUserController
// @Summary 激活用户
// @Description 通过用户id激活用户
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "用户id"
// @Success 200 {object} response.Response
// @Router /api/user/enable/{id} [get]
func (uc UserController) EnableUserController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := uc.Service.ActiveUserService(id)
	ctx.JSON(http.StatusOK, res)
}

// DisableUserController
// @Summary 激活用户
// @Description 通过用户id禁用用户
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "用户id"
// @Success 200 {object} response.Response
// @Router /api/user/enable/{id} [get]
func (uc UserController) DisableUserController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := uc.Service.DisableUserService(id)
	ctx.JSON(http.StatusOK, res)
}

// ResetPwdController
// @Summary 重置密码
// @Description 通过用户id重置登录密码为123456
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "用户id"
// @Success 200 {object} response.Response
// @Router /api/user/reset/{id} [get]
func (uc UserController) ResetPwdController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	zap.L().Info(ctx.Request.RequestURI, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := uc.Service.ResetPwdService(id)
	ctx.JSON(http.StatusOK, res)
}

// DeleteUserController
// @Summary 删除用户
// @Description 通过用户id删除用户
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer  true "用户id"
// @Success 200 {object} response.Response
// @Router /api/user/delete/{id} [delete]
func (uc UserController) DeleteUserController(ctx *gin.Context) {
	httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	msg := fmt.Sprintf("[url]=%s,[method]=%s", ctx.Request.URL.String(), ctx.Request.Method)
	zap.L().Info(msg, zap.String("request", string(httpRequest)))
	id, _ := com.StrTo(ctx.Param("id")).Int()
	res := uc.Service.DeleteUserService(id)
	ctx.JSON(http.StatusOK, res)
}
