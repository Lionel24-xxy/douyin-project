package user

import "errors"

const (
	MaxUsernameLength = 32
	MinPasswordLenth  = 8
)

type LoginAndRegisterResponse struct {
	UserId 		int64 	`json:"user_id"`
	Token  		string 	`json:"token"`
}

func isValidUser(name string, password string) error {
	if len(name) == 0 {
		return errors.New("用户名为空")
	}
	if len(name) > MaxUsernameLength {
		return errors.New("用户名长度超过32位")
	}
	if password == "" {
		return errors.New("密码为空")
	}
	if len(password) < MinPasswordLenth {
		return errors.New("密码少于8位")
	}
	return nil
}