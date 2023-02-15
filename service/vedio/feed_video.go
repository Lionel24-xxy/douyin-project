package vedio

import (
	"TikTok_Project/repository"

	"time"
)

//const maxVNum = 30

type feedVideo struct {
	Videos   []*repository.Video `json:"video_list,omitempty"`
	NextTime int64               `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*feedVideo, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}

type QueryFeedVideoListFlow struct {
	userId     int64
	latestTime time.Time
	videos     []*repository.Video
	nextTime   int64
	feedVideo  *feedVideo
}

func NewQueryFeedVideoListFlow(userId int64, latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{userId: userId, latestTime: latestTime}
}

func (q *QueryFeedVideoListFlow) Do() (*feedVideo, error) {
	//所有传入的参数不填也应该给他正常处理
	return q.feedVideo, nil
}
