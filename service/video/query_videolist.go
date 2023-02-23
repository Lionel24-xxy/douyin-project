package video

import (
	"TikTok_Project/repository"
	"errors"
)

type List struct {
	Videos []*repository.Video `json:"video_list,omitempty"`
}

func QueryPublishList(userid int64) (*List, error) {

	// 检查用户是否存在
	if !repository.InitUserDao().IsExistUserId(userid) {
		return nil, errors.New("用户不存在")
	}

	// 发布视频列表及作者信息打包返回
	// 获取视频列表
	var videoList []*repository.Video
	if err := repository.NewVideoDAO().QueryPublishListById(userid, &videoList); err != nil {
		return nil, err
	}
	// 获取作者信息
	var user repository.User
	if err := repository.InitUserDao().UserInfoById(userid, &user); err != nil {
		return nil, err
	}
	// 给每个发布视频填充作者信息
	for i := range videoList {
		videoList[i].Author = user
		videoList[i].IsFavorite = repository.GetVideoFavorState(userid, videoList[i].Id)
		// fmt.Printf("i: %v\n", i)
	}

	publishList := &List{
		Videos: videoList,
	}
	return publishList, nil
}
