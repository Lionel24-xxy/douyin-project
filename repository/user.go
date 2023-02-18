package repository

import (
	"errors"
	"sync"
)

type User struct {
	ID 					int64 		`json:"id" gorm:"id,omitempty;primaryKey;AUTO_INCREMENT"`
	Username 			string 		`json:"username" gorm:"username,omitempty"`
	Password 			string 		`json:"password" gorm:"size:200;notnull"`
	FollowCount 		int64 		`json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount 		int64       `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      		bool        `json:"is_follow" gorm:"is_follow,omitempty"`
	Avatar				string		`json:"avatar" gorm:"avatar,omitempty"`
	BackgroundImage		string		`json:"background_image" gorm:"background_image,omitempty"`
	Signature			string		`json:"signature" gorm:"signature,omitempty"`
	TotalFavorited		int64		`json:"total_favorited" gorm:"total_favorited,omitempty"`
	WorkCount			int64		`json:"work_count" gorm:"work_count,omitempty"`
	FavoriteCount		int64		`json:"favorite_count" gorm:"favorite_count,omitempty"`
	Relations       	[]*User 	`json:"-" gorm:"many2many:user_relations;association_jointable_foreignkey:follow_id"`    //用户之间的多对多
	Videos        		[]*Video    `json:"-"`                                    											 //用户与投稿视频的一对多
	Favorite   			[]*Video    `json:"-" gorm:"many2many:user_favorite;"` 	   											 //用户与点赞视频之间的多对多
	Comments      		[]*Comment  `json:"-"`                                     											 //用户与评论的一对多
}

// 单例模式
type UserDAO struct {
}

var (
	userDao 	*UserDAO
	userOnce	sync.Once
)

func InitUserDao() *UserDAO {
	userOnce.Do(func() {
		userDao = new(UserDAO)
	})
	return userDao
}

// register
func (u *UserDAO) IsExistName(name string) bool{
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

// User
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