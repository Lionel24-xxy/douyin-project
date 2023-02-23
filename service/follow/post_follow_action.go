package follow

import (
	"TikTok_Project/repository"
	"errors"
)

const (
	FOLLOW = 1
	CANCEL = 2
)

var (
	ErrIvdAct    = errors.New("未定义操作")
	ErrIvdFolUsr = errors.New("关注用户不存在")
)

func PostFollowAction(userId, userToId int64, actionType int) error {
	return NewPostFollowAction(userId, userToId, actionType).Do()
}

type PostFollowActionFlow struct {
	userId     int64
	userToId   int64
	actionType int
}

func NewPostFollowAction(userId, userToId int64, actionType int) *PostFollowActionFlow {
	return &PostFollowActionFlow{userId: userId, userToId: userToId, actionType: actionType}
}

func (p *PostFollowActionFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}
	if err = p.publish(); err != nil {
		return err
	}
	return nil
}

func (p *PostFollowActionFlow) checkNum() error {
	if !repository.InitUserDao().IsExistUserId(p.userToId) {
		return ErrIvdFolUsr
	}
	if p.actionType != FOLLOW && p.actionType != CANCEL {
		return ErrIvdAct
	}

	//自己不能关注自己
	if p.userId == p.userToId {
		return ErrIvdAct
	}
	return nil
}

func (p *PostFollowActionFlow) publish() error {
	userDAO := repository.InitUserDao()
	var err error
	switch p.actionType {
	case FOLLOW:
		err = userDAO.AddUserFollow(p.userId, p.userToId)
		repository.UpdateUserRelation(p.userId, p.userToId, true)
	case CANCEL:
		err = userDAO.CancelUserFollow(p.userId, p.userToId)
		repository.UpdateUserRelation(p.userId, p.userToId, true)
	default:
		return ErrIvdAct
	}
	return err
}
