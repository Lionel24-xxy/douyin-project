package video

import (
	"TikTok_Project/repository"
	"time"
)

const maxVNum int = 30

type FeedVideoList struct {
	VideoList []*repository.Video `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}

type QueryFeedVideoListFlow struct {
	userId     int64
	latestTime time.Time
	videos     []*repository.Video
	nextTime   int64
	feedVideo  *FeedVideoList
}

func NewQueryFeedVideoListFlow(userId int64, latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{userId: userId, latestTime: latestTime}
}

func (q *QueryFeedVideoListFlow) Do() (*FeedVideoList, error) {
	q.check()
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.feedVideo, nil
}
func (q *QueryFeedVideoListFlow) packData() error {
	q.feedVideo = &FeedVideoList{
		VideoList: q.videos,
		NextTime:  q.nextTime,
	}
	return nil
}

func (q *QueryFeedVideoListFlow) check() {
	//上层通过把userId置零，表示userId不存在或不需要
	if q.userId > 0 {
		//这里说明userId是有效的，可以定制性的做一些登录用户的专属视频推荐
	}

	if q.latestTime.IsZero() {
		q.latestTime = time.Now()
	}
}

func (q *QueryFeedVideoListFlow) prepareData() error {
	err := repository.NewVideoDAO().QueryVideoListByLimitAndTime(maxVNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	q.nextTime = time.Now().Unix() / 1e6
	return nil
}
