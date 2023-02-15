package router

import (
	"github.com/gin-gonic/gin"

	"TikTok_Project/controller/user"
	"TikTok_Project/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	uGroup := r.Group("douyin")
	{
		uGroup.POST("/user/register/", user.UserRegister)
		uGroup.POST("/user/login/", user.UserLogin)
		uGroup.GET("/user/", middleware.JWTMiddleWare(), user.UserInfo)
	}
	return r
}