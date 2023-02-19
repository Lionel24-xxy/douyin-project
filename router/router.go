package router

import (
	"TikTok_Project/controller/feed"
	"github.com/gin-gonic/gin"

	"TikTok_Project/controller/user"
	"TikTok_Project/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 视频及截图存放位置
	r.Static("static", "./static")

	uGroup := r.Group("douyin")
	{
		// 基础接口
		uGroup.POST("/user/register/", user.UserRegister)
		uGroup.POST("/user/login/", user.UserLogin)
		uGroup.GET("/user/", middleware.JWTMiddleWare(), user.UserInfo)
		uGroup.POST("/publish/action/", middleware.JWTMiddleWare(), feed.PublishVideoHandler)
		uGroup.GET("/publish/list/", middleware.JWTMiddleWare(), feed.PublishListHandler)
		// 互动接口

		// 社交接口

	}
	return r
}
