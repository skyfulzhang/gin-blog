package main

import (
	"gin-blog/dao"
	"gin-blog/routes"
	"gin-blog/utils"
	"os"
	"os/signal"
	"syscall"
)

/*
参考文档 swagger
https://wnanbei.github.io/post/gin-%E9%85%8D%E7%BD%AE-swagger-%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/
https://blog.csdn.net/qq_45100706/article/details/115481714
*/

// @title Gin博客项目接口文档
// @version 1.0
// @description 博客主要包含用户、文章和分类
// @license.name Apache 2.0
func main() {
	dao.InitMysql()
	utils.ZapLogger()
	routes.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	routes.HttpServerStop()
}
