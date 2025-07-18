package endpoint

import (
	"ginDemo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JSON 传 name
type CreateFolderJson struct {
	Name string `json:"name"`
}

func CreateFolder(c *gin.Context) {
	//还要传父级目录id,接路径里的值
	parentIdStr := c.Param("parentId")

	parentId, err := strconv.Atoi(parentIdStr)
	if err != nil || parentId < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	var req CreateFolderJson
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	folder, err := service.CreateFolderService.CreateFolder(c, req.Name, parentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, folder)
}
