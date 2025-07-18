package middleware

import (
	"ginDemo/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 用于校验文件权限
func FilePermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileIdStr := c.Param("file_id")
		fileId, err := strconv.Atoi(fileIdStr)

		if err != nil || fileId < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "无效的file_id参数",
			})
			c.Abort()
			return
		}
		if fileId == 0 {
			c.Next()
			return
		}
		//目前有 登录的userId和fileId，先用fileId查对应flie的userid,在比较二者。
		userId := c.GetInt("user_id")

		file, err := dao.Q.File.WithContext(c.Request.Context()).
			Where(dao.File.ID.Eq(int32(fileId))).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"code": 404,
					"msg":  "文件不存在",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "数据库错误",
			})
			c.Abort()
			return
		}
		//二者对比
		if file.UserID != int32(userId) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "无权限访问",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
