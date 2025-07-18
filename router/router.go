package router

import (
	"ginDemo/endpoint"
	"ginDemo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建路由
	r := gin.Default()
	//不需要校验
	r.POST("/login", endpoint.Login)

	auth := r.Group("/api", middleware.Auth())

	//创建新用户(需要登录校验)
	auth.POST("/user", endpoint.CreateUser)

	//登录加文件双重校验的路由组
	authFile := auth.Group("/api/file", middleware.Auth(), middleware.FilePermissionMiddleWare())

	//创建文件夹
	authFile.POST("/:file_id/new", endpoint.CreateFolder)

	//上传文件(file_id为父级文件夹id)
	authFile.POST("/:file_id/upload", endpoint.UploadFile)
	return r

}
