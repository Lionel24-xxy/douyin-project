package follow

import (
	"TikTok_Project/service/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowerListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*user.FollowerList
}

func QueryFollowerHandler(c *gin.Context) {
	NewProxyQueryFollowerHandler(c).Do()
}

type ProxyQueryFollowerHandler struct {
	*gin.Context

	userId int64

	*user.FollowerList
}

func NewProxyQueryFollowerHandler(context *gin.Context) *ProxyQueryFollowerHandler {
	return &ProxyQueryFollowerHandler{Context: context}
}

func (p *ProxyQueryFollowerHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, user.UserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("失败")
		}
		return
	}
	p.SendOk("成功")
}

func (p *ProxyQueryFollowerHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowerHandler) prepareData() error {
	list, err := user.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

func (p *ProxyQueryFollowerHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyQueryFollowerHandler) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		StatusCode:   0,
		StatusMsg:    msg,
		FollowerList: p.FollowerList,
	})
}
