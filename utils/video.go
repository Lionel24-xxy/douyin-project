package utils

import (
	"TikTok_Project/repository"
	"errors"
	"fmt"
	"time"
)

var (
	defaultIP   = "192.168.128.122"
	defaultPort = 8080
)

// NewFileName 根据UserID+雪花算法生成的id连接成videoName
func NewFileName(userID int64) string {
	node, _ := NewWorker(1)
	randomID := node.NextId()
	return fmt.Sprintf("%d-%d", userID, randomID)
}

// GetFileUrl 通过filename，将它改为url形式
func GetFileUrl(filename string) string {
	base := fmt.Sprintf("http://%s:%d/static/%s", defaultIP, defaultPort, filename)
	return base
}

// FillVideoListFields 填充每个视频的作者信息（因为作者与视频的一对多关系，数据库中存下的是作者的id
// 当userId>0时，我们判断当前为登录状态，其余情况为未登录状态，则不需要填充IsFavorite字段
func FillVideoListFields(userId int64, videos *[]*repository.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	dao := repository.InitUserDao()

	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	//添加作者信息，以及is_follow状态
	for i := 0; i < size; i++ {
		var userInfo repository.User
		err := dao.UserInfoById((*videos)[i].UserId, &userInfo)
		if err != nil {
			continue
		}
		userInfo.IsFollow = repository.GetUserRelation(userId, userInfo.ID) //根据cache更新是否被点赞
		(*videos)[i].Author = userInfo
		//填充有登录信息的点赞状态
		if userId > 0 {
			(*videos)[i].IsFavorite = repository.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}
