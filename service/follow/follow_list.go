package follow

import (
	"TikTok_Project/repository"
	"errors"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

type FollowList struct {
	UserList []*repository.User `json:"user_list"`
}

type QueryFollowListFlow struct {
	userId   int64
	userList []*repository.User
	*FollowList
}

func NewQueryFollowListFlow(userId int64) *QueryFollowListFlow {
	return &QueryFollowListFlow{userId: userId}
}

func QueryFollowList(userId int64) (*FollowList, error) {
	return NewQueryFollowListFlow(userId).DO()
}

func (q *QueryFollowListFlow) DO() (*FollowList, error) {
	var err error
	if err = q.checkNum(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}
	return q.FollowList, nil
}

func (q *QueryFollowListFlow) checkNum() error {
	if !repository.InitUserDao().IsExistUserId(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFollowListFlow) prepareData() error {
	var userList []*repository.User
	err := repository.InitUserDao().GetFollowListByUserId(q.userId, &userList)
	if err != nil {
		return err
	}
	for i, _ := range userList {
		userList[i].IsFollow = true //当前用户的关注列表，故isFollow定为true
	}
	q.userList = userList
	return nil
}

func (q *QueryFollowListFlow) packData() error {
	q.FollowList = &FollowList{UserList: q.userList}
	return nil
}
