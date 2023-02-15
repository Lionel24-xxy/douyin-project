package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"TikTok_Project/utils"
)

// JWTMiddleWare 鉴权中间件，鉴权并设置user_id
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := utils.ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 5, 
				"status_msg": "token不正确",
			})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 5, 
				"status_msg": "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.UserId)
		c.Next()
	}
}