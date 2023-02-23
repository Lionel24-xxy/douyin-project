package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"TikTok_Project/repository"
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


// RateMiddleware 中间件
func RateMiddleware(c *gin.Context) {
	//以Pipeline的方式操作事务
	pipe := repository.RDB.TxPipeline()

	// 5 秒刷新key为IP(c.ClientIP())的r值为0
	err := pipe.SetNX(repository.CTX, c.ClientIP(), 0, 10 * time.Second).Err()
	if err != nil {
		log.Printf("redis刷新错误" + err.Error())
	}
	// 每次访问，这个IP的对应的值加一
	pipe.Incr(repository.CTX, c.ClientIP())
	// 提交事务
	_, _ = pipe.Exec(repository.CTX)

	// 获取IP访问的次数
	var val int
	val, err = repository.RDB.Get(repository.CTX, c.ClientIP()).Int()
	if err != nil {
		log.Printf("redis刷新错误" + err.Error())
	}
	// 如果10秒内大于50次
	if val > 50 {
		c.Abort()
		c.JSON(http.StatusOK, gin.H{
			"status_code" : -1,
			"status_msg" : "访问过于频繁",
		})
	} else {
		c.Next()
	}
}