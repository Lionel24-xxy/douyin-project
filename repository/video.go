package repository

import "time"

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