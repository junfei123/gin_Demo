package endpoint

import (
	"errors"
	"ginDemo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 前端传username和password(JSON)
type CreateUserRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func CreateUser(c *gin.Context) {
	//校验是否为管理员
	userName := c.GetString("user_name")
	if userName != "admin" {
		c.JSON(403, gin.H{
			"code": 403,
			"msg":  "无权限创建用户"})
		return
	}
	//JSON->结构体
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误"})
		return
	}

	if err := service.CreateUserService.CreateUser(c.Request.Context(), req.Username, req.Password); err != nil {
		//唯一索引冲突
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "用户已存在"})
			return
		}

		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "创建用户失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建用户成功",
	})

}
