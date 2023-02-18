package user

import (
	"TikTok_Project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "解析userId出错"})
		return
	}
	userId, _ := rawId.(int64)
	//fmt.Printf("userId: %v\n", userId)
	var user repository.User
	userInfoDao := repository.InitUserDao()
	if err := userInfoDao.UserInfoById(userId, &user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 2,
			"status_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "Access User Information success!",
		"user": gin.H{
			"id":             	user.ID,
			"name":           	user.Username,
			"follow_count":   	user.FollowCount,
			"follower_count": 	user.FollowerCount,
			"is_follow":      	user.IsFollow,
			"avatar": 		  	user.Avatar,
            "background_image": user.BackgroundImage,
            "signature": 		user.Signature,
            "total_favorited": 	user.TotalFavorited,
            "work_count": 		user.WorkCount,
            "favorite_count": 	user.FavoriteCount,
		}})
}

