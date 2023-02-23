package comment

import (
	"TikTok_Project/service/comment"
	"TikTok_Project/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublishCommentResponse struct {
	StatueCode		int			`json:"status_code"`
	StatueMsg		string		`json:"status_msg"`
	*comment.CommentResponse
}


type ProxyPostCommentHandler struct {
	videoId     int64
	userId      int64
	commentId   int64
	actionType  int64
	commentText string
}

func PublishCommentHandler(c *gin.Context) {
	var p = &ProxyPostCommentHandler{}
	if err := p.prepareParse(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	comment, err := comment.PublishComment(p.userId, p.videoId, p.actionType, p.commentText, p.commentId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PublishCommentResponse{
		StatueCode: 0,
		StatueMsg: "评论操作成功！",
		CommentResponse: comment,
	})
}

func (p *ProxyPostCommentHandler) prepareParse(c *gin.Context) error {
	rawId, _ := c.Get("user_id")
	userId, ok := rawId.(int64)
	//fmt.Printf("userId: %v\n", userId)
	if !ok {
		return errors.New("解析userId出错") 
	}
	p.userId = userId
	
	rawVideoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	rawActionType := c.Query("action_type")
	actiontype, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		return err
	}
	// 判断是发布还是删除
	switch actiontype {
	case comment.CREATE_COMMENT: // 发布评论
		rawText := c.Query("comment_text")
		p.commentText = p.sensitiveCheck(rawText)

	case comment.DELETE_COMMENT: // 删除评论
		p.commentId, err = strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("未定义的行为：%d", actiontype)
	}
	p.actionType = actiontype
	return nil
}

// 敏感词检测及替换
func (p *ProxyPostCommentHandler) sensitiveCheck(text string) string {
	
	isContain := utils.SensitiveWordCheck(text, int(p.userId))
	if isContain {
		replaceText := utils.SensitiveWordReplace(text)
		return replaceText
	}
	return text
}