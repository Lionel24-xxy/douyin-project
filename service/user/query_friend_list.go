package user

import (
	"TikTok_Project/repository"
)

type FriendList struct {
	UserList []*repository.User `json:"user_list"`
}
type QueryFriendListFlow struct {
	userId int64

	userList []*repository.User

	*FriendList
}

func QueryFriendList(userId int64) (*FriendList, error) {
	return NewQueryFriendListFlow(userId).Do()
}

func NewQueryFriendListFlow(userId int64) *QueryFriendListFlow {
	return &QueryFriendListFlow{userId: userId}
}

func (q *QueryFriendListFlow) Do() (*FriendList, error) {
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
	return q.FriendList, nil
}

func (q *QueryFriendListFlow) checkNum() error {
	if !repository.InitUserDao().IsExistUserId(q.userId) {
		return UserNotExist
	}
	return nil
}

func (q *QueryFriendListFlow) prepareData() error {

	err := repository.InitUserDao().GetFriendListByUserId(q.userId, &q.userList)
	if err != nil {
		return err
	}

	for _, v := range q.userList {
		//获取用户关注列表接口
		v.IsFollow = repository.GetUserRelation(q.userId, v.ID)
	}
	return nil
}

func (q *QueryFriendListFlow) packData() error {
	q.FriendList = &FriendList{UserList: q.userList}
	return nil
}
