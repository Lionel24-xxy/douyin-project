package comment

import (
	"TikTok_Project/repository"
	"TikTok_Project/utils"
	"errors"

	"fmt"
)

const (
	CREATE_COMMENT int64 = 1
	DELETE_COMMENT int64 = 2
)

type CommentResponse struct {
	Comments *repository.Comment `json:"comment"`
}

func PublishComment(userId int64, videoId int64, actiontype int64, commentText string, commentId int64) (*CommentResponse, error) {
	/// 参数检查
	// 检查用户是否存在
	if !repository.InitUserDao().IsExistUserId(userId) {
		return nil, fmt.Errorf("用户%d不存在", userId)
	}
	// 检查视频是否存在
	if !repository.NewVideoDAO().IsExistVideoById(videoId) {
		return nil, fmt.Errorf("视频%d不存在", videoId)
	}
	// 判断actiontype
	if actiontype != 1 && actiontype != 2 {
		return nil, fmt.Errorf("未定义行为：%d", actiontype)
	}

	var comment repository.Comment
	switch actiontype {
	case 1:
		/// 发布评论
		comment = repository.Comment{
			UserId:  userId,
			VideoId: videoId,
			Content: commentText,
		}
		err := repository.InitCommentDAO().AddComment(&comment)
		if err != nil {
			return nil, err
		}
	case 2:
		/// 删除评论
		// 根据commentId找到评论
		err := repository.InitCommentDAO().QueryCommentById(commentId, &comment)
		if err != nil {
			return nil, err
		}
		// 删除comment
		err = repository.InitCommentDAO().DeleteComment(commentId, videoId)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("未定义的操作")
	}

	/// 返回结果
	// 添加user信息
	var user repository.User
	if err := repository.InitUserDao().UserInfoById(userId, &user); err != nil {
		return nil, err
	}
	comment.User = user
	// 添加发布日期
	if err := utils.FillCommentFields(&comment); err != nil {
		return nil, err
	}

	var commentResponse = &CommentResponse{Comments: &comment}
	return commentResponse, nil
}
