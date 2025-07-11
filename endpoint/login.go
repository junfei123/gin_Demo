package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接受JSON串
type LoginJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginJson LoginJson
	if err := c.ShouldBindJSON(&loginJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

}
