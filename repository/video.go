package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Id            int64      `json:"id,omitempty"`
	UserId        int64      `json:"-"`
	Author        User       `json:"author,omitempty" gorm:"-"` //这里应该是作者对视频的一对多的关系，而不是视频对作者，故gorm不能存他，但json需要返回它
	PlayUrl       string     `json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	IsFavorite    bool       `json:"is_favorite,omitempty"`
	Title         string     `json:"title,omitempty"`
	Users         []*User    `json:"-" gorm:"many2many:user_favorite;"`
	Comments      []*Comment `json:"-"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
}

type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}

// QueryVideoListByLimitAndTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func (v *VideoDAO) QueryVideoListByLimitAndTime(limit int, latestTime time.Time, videoList *[]*Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}
	return DB.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at ASC").Limit(limit).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title", "created_at", "updated_at"}).
		Find(videoList).Error
}

// AddVideo 添加视频
// 注意：由于视频和userinfo有多对一的关系，所以传入的Video参数一定要进行id的映射处理！
func (v *VideoDAO) AddVideo(video *Video) error {
	if video == nil {
		return errors.New("AddVideo video 空指针")
	}
	return DB.Create(video).Error
}
// 更新发布视频数
func (v *VideoDAO) UpdateWorkCount(UserId int64) error {
	var user User
	DB.First(&user, "id = ?", UserId)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	err := DB.Model(&user).Update("work_count", gorm.Expr("work_count+1")).Error
	return err
}
