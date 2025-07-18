package middleware

import (
	"ginDemo/dao"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		sid, err := c.Cookie("sid")
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "未登录",
			})
			// 拦截
			c.Abort()
			return

		}

		// 验证sid
		session, err := dao.Q.Session.WithContext(c.Request.Context()).
			Where(dao.Session.SessionID.Eq(sid)).First()
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "无效的sid",
			})
			c.Abort()
			return
		}
		//验证userid(可能存在sid，但userid不存在，用户被干了)
		user, err := dao.Q.User.WithContext(c.Request.Context()).
			Where(dao.User.ID.Eq(session.UserID)).First()
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "无效的userid",
			})
			c.Abort()
			return
		}
		c.Set("user_id", int(user.ID))
		c.Set("user_name", user.Name)
		c.Next()
	}
}
