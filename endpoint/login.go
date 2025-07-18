package endpoint

import (
	"ginDemo/dao"
	"ginDemo/model"
	"ginDemo/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// 账号密码匹配
	uid, err := service.LoginService.Login(c.Request.Context(), loginJson.Username, loginJson.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}

	//会话id,存到session中
	sessionId := uuid.New().String()

	dao.Q.Session.WithContext(c.Request.Context()).Create(&model.Session{
		UserID:    uid,
		SessionID: sessionId,
		Ctime:     int32(time.Now().Unix()),
		Etime:     int32(time.Now().Add(time.Hour * 1).Unix()),
	})
	c.SetCookie("sid", sessionId, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})

}
