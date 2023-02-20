package user

import (
	"TikTok_Project/repository"
	"errors"
)

type FollowerList struct {
	UserList []*repository.User `json:"user_list"`
}

type QueryFollowerListFlow struct {
	userId int64

	userList []*repository.User

	*FollowerList
}

var (
	UserNotExist = errors.New("用户不存在或已注销")
)

func QueryFollowerList(userId int64) (*FollowerList, error) {
	return NewQueryFollowerListFlow(userId).Do()
}

func NewQueryFollowerListFlow(userId int64) *QueryFollowerListFlow {
	return &QueryFollowerListFlow{userId: userId}
}

func (q *QueryFollowerListFlow) Do() (*FollowerList, error) {
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
	return q.FollowerList, nil
}

func (q *QueryFollowerListFlow) checkNum() error {
	if !repository.InitUserDao().IsExistUserId(q.userId) {
		return UserNotExist
	}
	return nil
}

func (q *QueryFollowerListFlow) prepareData() error {

	err := repository.InitUserDao().GetFollowerListByUserId(q.userId, &q.userList)
	if err != nil {
		return err
	}
	//填充is_follow字段

	//待做
	//for _, v := range q.userList {
	//	//获取用户关注列表接口
	//	//v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(q.userId, v.Id)
	//}
	return nil
}

func (q *QueryFollowerListFlow) packData() error {
	q.FollowerList = &FollowerList{UserList: q.userList}

	return nil
}
