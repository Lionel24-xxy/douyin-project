package feed

import (
	"TikTok_Project/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ProxyFeedVideoList struct {
	*gin.Context
}

//func ProcessFeedVideoList(c *gin.Context) {
//	p := NewProxyFeedVideoList(c)
//	//无登录状态
//	if !ok {
//		err := p.DoNoLog()
//		if err != nil {
//			p.FeedVideoListError(err.Error())
//		}
//		return
//	}
//
//}

func NewProxyFeedVideoList(c *gin.Context) *ProxyFeedVideoList {
	return &ProxyFeedVideoList{Context: c}
}

// DoNoLog 未登录的视频流推送处理
func (p *ProxyFeedVideoList) DoNoLog() error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	videoList, err := video.QueryFeedVideoList(0, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

func (p *ProxyFeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyFeedVideoList) FeedVideoListOk(videoList *video.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		StatusCode: 0,
	},
	)
}
