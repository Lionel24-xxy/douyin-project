package comment

import (
	"TikTok_Project/repository"
	"TikTok_Project/utils"
	"errors"
	"fmt"
)

type CommentList struct {
	Comments []*repository.Comment `json:"comment_list"`
}

func QueryCommentList(userId, videoId int64) (*CommentList, error) {
	// 检查数据是否存在
	if !repository.InitUserDao().IsExistUserId(userId) {
		return nil, fmt.Errorf("用户%d处于登出状态", userId)
	}
	if !repository.NewVideoDAO().IsExistVideoById(videoId) {
		return nil, fmt.Errorf("视频%d不存在或已经被删除", videoId)
	}

	// 查询视频评论列表
	var commentList []*repository.Comment
	err := repository.InitCommentDAO().QueryCommentListByVideoId(videoId, &commentList)
	if err != nil {
		return nil, err
	}
	// 根据前端的要求填充正确的时间格式
	err = utils.FillCommentListFields(&commentList)
	if err != nil {
		return nil, errors.New("暂时还没有人评论")
	}
	// 添加评论的作者信息
	var user repository.User
	err = repository.InitUserDao().UserInfoById(userId, &user)
	if err != nil {
		return nil, err
	}
	for i := range commentList {
		commentList[i].User = user
	}

	// 返回结果
	response := &CommentList{Comments: commentList}
	return response, nil
}
