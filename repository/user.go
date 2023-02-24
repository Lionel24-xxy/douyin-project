package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var (
	EmptyUserList = errors.New("用户粉丝列表为空")
)

type User struct {
	ID              int64      `json:"id" gorm:"id,omitempty;primaryKey;AUTO_INCREMENT"`
	Username        string     `json:"username" gorm:"username,omitempty"`
	Password        string     `json:"password" gorm:"size:200;notnull"`
	FollowCount     int64      `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount   int64      `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow        bool       `json:"is_follow" gorm:"is_follow,omitempty"`
	Avatar          string     `json:"avatar" gorm:"avatar,omitempty"`
	BackgroundImage string     `json:"background_image" gorm:"background_image,omitempty"`
	Signature       string     `json:"signature" gorm:"signature,omitempty"`
	TotalFavorited  int64      `json:"total_favorited" gorm:"total_favorited,omitempty"`
	WorkCount       int64      `json:"work_count" gorm:"work_count,omitempty"`
	FavoriteCount   int64      `json:"favorite_count" gorm:"favorite_count,omitempty"`
	Relations       []*User    `json:"-" gorm:"many2many:user_relations;association_jointable_foreignkey:follow_id"` //用户之间的多对多
	Videos          []*Video   `json:"-"`                                                                            //用户与投稿视频的一对多
	Favorite        []*Video   `json:"-" gorm:"many2many:user_favorite;"`                                            //用户与点赞视频之间的多对多
	Comments        []*Comment `json:"-"`                                                                            //用户与评论的一对多
}

// 单例模式
type UserDAO struct {
}

var (
	userDao  *UserDAO
	userOnce sync.Once
)

func InitUserDao() *UserDAO {
	userOnce.Do(func() {
		userDao = new(UserDAO)
	})
	return userDao
}

// register
func (u *UserDAO) IsExistName(name string) bool {
	var user User
	DB.Find(&user, "username = ?", name)
	return user.ID != 0
}

func (u *UserDAO) UserRegister(user *User) error {
	err := DB.Create(&user).Error
	return err
}

// login
func (u *UserDAO) UserLoginVerify(name string, password string, user *User) bool {
	DB.First(&user, "username = ? and password = ?", name, password)

	return user.ID == 0
}

// UserInfo
func (u *UserDAO) UserInfoById(userId int64, user *User) error {
	if user == nil {
		return errors.New("空指针错误")
	}

	DB.Where("id=?", userId).Select([]string{"id", "username", "follow_count", "follower_count", "is_follow", "avatar", "background_image", "signature", "total_favorited", "work_count", "favorite_count"}).First(user)

	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	return nil
}

// 用户是否存在
func (u *UserDAO) IsExistUserId(userid int64) bool {
	var userinfo User
	if err := DB.Where("id=?", userid).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.ID == 0 {
		return false
	}
	return true
}

func (u *UserDAO) GetFollowerListByUserId(userId int64, userList *[]*User) error {
	if userList == nil {
		return errors.New("空指针错误")
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, users u WHERE r.follow_id = ? AND r.user_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}

	if len(*userList) == 0 || (*userList)[0].ID == 0 {
		return EmptyUserList
	}

	return nil
}

func (u *UserDAO) AddUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error { //事务以块的形式开始一个事务，返回错误会回滚，否则就提交。
		if err := tx.Exec("UPDATE users SET follow_count=follow_count+1 WHERE id=?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE users SET follower_count=follower_count+1 WHERE id=?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations`  (`user_id`,`follow_id`) VALUES (?,?)", userId, userToId).Error; err != nil { //连接表user_relations
			return err
		}
		return nil
	})
}

func (u *UserDAO) CancelUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE users SET follow_count=follow_count-1 WHERE id = ? AND follow_count>0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE users SET follower_count=follower_count-1 WHERE id = ? AND follower_count>0", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_id=? AND follow_id=?", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserDAO) GetFollowListByUserId(userId int64, userList *[]*User) error {
	if userList == nil {
		return errors.New("空指针错误")
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, users u WHERE r.user_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].ID == 0 {
		return errors.New("用户列表为空")
	}
	return nil
}

func (u *UserDAO) GetFriendListByUserId(userId int64, userList *[]*User) error {
	if userList == nil {
		return errors.New("空指针错误")
	}
	var err error
	//if err = DB.Raw("SELECT u.* FROM user_relations r,users u WHERE r.follow_id = ? AND r.user_id = u.id", userId).Scan(userList).Error; err != nil {
	if err = DB.Raw("SELECT t1.* FROM (SELECT * FROM user_relations WHERE `user_id` = ?) AS t1 INNER JOIN (SELECT * FROM user_relations WHERE `follow_id` = ?) AS t2 ON t1.follow_id = t2.user_id", userId).Scan(userList).Error; err != nil {
		return err
	}
	return nil
}
