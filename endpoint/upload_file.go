package endpoint

import (
	"ginDemo/service"
	"ginDemo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	//校验
	parentIdStr := c.Param("parent_id")
	parentId, err := strconv.Atoi(parentIdStr)

	if err != nil || parentId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
	}

	//1. 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "上传文件错误",
		})
		return
	}
	//2. 获取文件名
	fileName := file.Filename
	userId := c.GetInt("user_id")
	//重名检测
	fileName = utils.AutoRenameFile(c.Request.Context(), int32(userId), fileName, int32(parentId))

	//3. 保存文件调service
	uploadFile, err := service.UploadFileService.UploadFile(c, file, fileName, parentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "上传文件错误",
		})
		return
	}

	//4. 返回文件基本信息
	c.JSON(http.StatusOK, uploadFile)

}
