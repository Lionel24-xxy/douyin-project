package video

import (
	"TikTok_Project/repository"
	"time"
)

// MaxVideoNum 每次最多返回的视频流数量
const (
	MaxVideoNum = 30
)

type FeedVideoList struct {
	Videos   []*repository.Video `json:"video_list,omitempty"`
	NextTime int64               `json:"next_time,omitempty"`
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
	//上层通过把userId置零，表示userId不存在或不需要
	if q.userId > 0 {
		//这里说明userId是有效的，可以定制性的做一些登录用户的专属视频推荐
	}

	if q.latestTime.IsZero() {
		q.latestTime = time.Now()
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.feedVideo, nil
}

func (q *QueryFeedVideoListFlow) prepareData() error {
	err := repository.NewVideoDAO().QueryVideoListByLimitAndTime(MaxVideoNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	//还有要补充的
	q.nextTime = time.Now().Unix() / 1e6
	return nil
}

func (q *QueryFeedVideoListFlow) packData() error {
	q.feedVideo = &FeedVideoList{
		Videos:   q.videos,
		NextTime: q.nextTime,
	}
	return nil
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}
