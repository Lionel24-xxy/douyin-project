package feed

import (
	"TikTok_Project/service/video"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublishListResponse struct {
	StatusCode 	int				`json:"status_code"`
	StatusMsg	string			`json:"status_msg"`
				*video.List
}

func PublishListHandler(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "解析userId出错"})
		return
	}
	userId, _ := rawId.(int64)

	publishList, err := video.QueryPublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		StatusCode: 0,
		StatusMsg: "读取发布视频列表成功！",
		List: publishList,
	})
}