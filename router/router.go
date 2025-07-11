package router

import (
	"ginDemo/endpoint"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建路由
	r := gin.Default()

	r.POST("/login", endpoint.Login)
	return r
}
