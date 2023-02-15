package user

import (
	"TikTok_Project/repository"
	"TikTok_Project/utils"
	"errors"

	"github.com/gin-gonic/gin"
)


func UserRegister(c *gin.Context, name string, password string) (*LoginAndRegisterResponse, int, error) {

	if err := isValidUser(name, password); err != nil {
		return nil, 1, err
	}

	// 判断用户名是否存在
	userExistDao := repository.InitUserDao()
	if userExistDao.IsExistName(name) {
		return nil, 2, errors.New("用户名已存在")
	}

	var user repository.User
	user.Username = name
	user.Password = password

	// 密码加密
	user.Password = utils.SHA1(password)

	// 数据库更新用户数据
	userUpdateDao := repository.InitUserDao()
	err := userUpdateDao.UserRegister(&user)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 3, "status_msg": err.Error()})
		return nil, 3, err
	}

	// 获取 token
	token, err := utils.GenToken(user)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 4, "status_msg": err.Error()})
		return nil, 4, err
	}
	
	respose := &LoginAndRegisterResponse{
		UserId: user.ID,
		Token: token,
	}
	return respose, 0, nil
}