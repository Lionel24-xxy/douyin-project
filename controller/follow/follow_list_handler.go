package follow

import (
	"TikTok_Project/service/follow"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*follow.FollowList
}

type ProxyQueryFollowList struct {
	*gin.Context
	UserId int64
	*follow.FollowList
}

func NewProxyQueryFollowList(ctx *gin.Context) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Context: ctx}
}

func QueryFollowListHandler(ctx *gin.Context) {
	NewProxyQueryFollowList(ctx).DO()
}

func (p *ProxyQueryFollowList) SendErr(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyQueryFollowList) SendOK(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  msg,
		FollowList: p.FollowList,
	})
}

func (p *ProxyQueryFollowList) DO() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendErr(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendErr(err.Error())
		return
	}
	p.SendOK("请求成功")
}

func (p *ProxyQueryFollowList) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("UserId解析错误")
	}
	p.UserId = userId
	return nil
}

func (p *ProxyQueryFollowList) prepareData() error {
	list, err := follow.QueryFollowList(p.UserId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}
