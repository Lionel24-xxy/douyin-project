package comment

import (
	"TikTok_Project/service/comment"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProxyCommentListHandler struct {
	videoId int64
	userId  int64
}

type CommentListResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	*comment.CommentList
}

func CommentListHandler(c *gin.Context) {
	var p ProxyCommentListHandler
	err := p.prepareParse(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "statue_msg": err})
		return
	}

	commentList, err := comment.QueryCommentList(p.userId, p.videoId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 2, "statue_msg": err})
		return
	}

	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功！",
		CommentList: commentList,
	})
}

func (p *ProxyCommentListHandler) prepareParse(c *gin.Context) error {
	rawUserId, _ := c.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	rawVideoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}
