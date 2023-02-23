package router

import (
	"TikTok_Project/controller/feed"
	"TikTok_Project/controller/follow"
	"TikTok_Project/controller/video"

	"github.com/gin-gonic/gin"

	"TikTok_Project/controller/comment"
	"TikTok_Project/controller/user"
	"TikTok_Project/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// 视频及截图存放位置
	r.Static("static", "./static")

	uGroup := r.Group("douyin", middleware.RateMiddleware)
	{
		// 基础接口
		uGroup.GET("/feed/", feed.FeedVideoListHandler)
		uGroup.POST("/user/register/", user.UserRegister)
		uGroup.POST("/user/login/", user.UserLogin)
		uGroup.GET("/user/", middleware.JWTMiddleWare(), user.UserInfo)
		uGroup.POST("/publish/action/", middleware.JWTMiddleWare(), feed.PublishVideoHandler)
		uGroup.GET("/publish/list/", middleware.JWTMiddleWare(), feed.PublishListHandler)

		// 互动接口
		uGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), video.PostFavorHandler)
		uGroup.GET("/favorite/list/", middleware.JWTMiddleWare(), video.QueryFavorVideoListHandler)
		uGroup.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PublishCommentHandler)
		uGroup.GET("/comment/list/", middleware.JWTMiddleWare(), comment.CommentListHandler)

		// 社交接口
		uGroup.POST("/relation/action/", middleware.JWTMiddleWare(), follow.PostFollowActionHandler)
		uGroup.GET("/relation/follow/list/", middleware.NoAuthToGetUserId(), follow.QueryFollowListHandler)
		uGroup.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), follow.QueryFollowerHandler)
	}
	return r
}
