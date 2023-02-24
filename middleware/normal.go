package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NoAuthToGetUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawId := ctx.Query("user_id")
		if rawId == "" {
			rawId = ctx.PostForm("user_id")
		}
		//用户不存在
		if rawId == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"statuscode": 401,
				"statusmsg":  "用户不存在",
			})
			ctx.Abort() //阻止执行
			return
		}
		userId, err := strconv.ParseInt(rawId, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statuscode": 401,
				"statusmsg":  "用户不存在",
			})
			ctx.Abort() //阻止执行
			return
		}
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
