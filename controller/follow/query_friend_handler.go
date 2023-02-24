package follow

import (
	"TikTok_Project/service/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FriendListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*user.FriendList
}

func QueryFriendHandler(c *gin.Context) {
	NewProxyQueryFollowerHandler(c).Do()
}

type ProxyQueryFriendHandler struct {
	*gin.Context
	userId int64
	*user.FriendList
}

func NewProxyQueryFriendHandler(context *gin.Context) *ProxyQueryFriendHandler {
	return &ProxyQueryFriendHandler{Context: context}
}

func (p *ProxyQueryFriendHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, user.UserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("准备数据出错")
		}
		return
	}
	p.SendOk("成功")
}

func (p *ProxyQueryFriendHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFriendHandler) prepareData() error {
	list, err := user.QueryFriendList(p.userId)
	if err != nil {
		return err
	}
	p.FriendList = list
	return nil
}

func (p *ProxyQueryFriendHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyQueryFriendHandler) SendOk(msg string) {
	p.JSON(http.StatusOK, FriendListResponse{
		StatusCode: 0,
		StatusMsg:  msg,
		FriendList: p.FriendList,
	})

}
