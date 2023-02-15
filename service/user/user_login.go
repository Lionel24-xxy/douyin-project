package user

import (
	"TikTok_Project/repository"
	"TikTok_Project/utils"
	"errors"
)



func UserLogin(name string, password string) (*LoginAndRegisterResponse, int, error) {
	// 验证用户名和密码是否合法
	if err := IsValidUser(name, password); err != nil {
		return nil, 1, err
	}

	var user repository.User

	// 验证用户名和密码
	password = utils.SHA1(password)
	userDao := repository.InitUserDao()
	if userDao.UserLoginVerify(name, password, &user) {
		//c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": "name is not exist or password is wrong"})
		return nil, 2, errors.New("用户不存在，账号或密码出错")
	}

	// 获取 token
	token, err := utils.GenToken(user)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 4, "status_msg": err.Error()})
		return nil, 4, err
	}

	// 返回结果
	response := &LoginAndRegisterResponse{
		UserId: user.ID,
		Token: token,
	}
	return response, 0, nil
}