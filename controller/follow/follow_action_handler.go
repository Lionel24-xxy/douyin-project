package follow

import (
	"TikTok_Project/service/follow"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FollowResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func PostFollowActionHandler(c *gin.Context) {
	NewProxyPostFollowAction(c).Do()
}

type ProxyPostFollowAction struct {
	*gin.Context

	userId     int64
	followId   int64
	actionType int
}

func NewProxyPostFollowAction(ctx *gin.Context) *ProxyPostFollowAction {
	return &ProxyPostFollowAction{Context: ctx}
}

func (p *ProxyPostFollowAction) Do() {
	var err error
	if err = p.prepareNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.startAction(); err != nil {
		//当错误在model层发生的，那么就是重复键值的插入
		if errors.Is(err, follow.ErrIvdAct) || errors.Is(err, follow.ErrIvdFolUsr) {
			p.SendError(err.Error())
		} else {
			p.SendError("请勿重复关注")
		}
		return
	}
	p.SendOK("操作成功")
}

func (p *ProxyPostFollowAction) prepareNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("UserId解析出错")
	}
	p.userId = userId

	//解析需要关注的id
	followId := p.Query("to_user_id")
	followidInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		return err
	}

	p.followId = followidInt

	//解析action_type
	actionType := p.Query("action_type")
	actiontypeInt, Err := strconv.ParseInt(actionType, 10, 64)
	if Err != nil {
		return Err
	}
	p.actionType = int(actiontypeInt)
	return nil
}

func (p *ProxyPostFollowAction) startAction() error {
	err := follow.PostFollowAction(p.userId, p.followId, p.actionType)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (p *ProxyPostFollowAction) SendError(msg string) {
	p.JSON(http.StatusOK, FollowResponse{StatusCode: 1, StatusMsg: msg})
}

func (p *ProxyPostFollowAction) SendOK(msg string) {
	p.JSON(http.StatusOK, FollowResponse{StatusCode: 0, StatusMsg: msg})
}
