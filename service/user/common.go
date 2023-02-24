package user

import (
	"github.com/dlclark/regexp2"

	"errors"
)

const (
	MaxUsernameLength = 32
	MinPasswordLength  = 8
)

type LoginAndRegisterResponse struct {
	UserId 		int64 	`json:"user_id"`
	Token  		string 	`json:"token"`
}

func IsValidUser(name string, password string) error {
	if len(name) == 0 {
		return errors.New("用户名为空")
	}
	if len(name) > MaxUsernameLength {
		return errors.New("用户名长度超过32位")
	}
	if password == "" {
		return errors.New("密码为空")
	}
	if len(password) < MinPasswordLength {
		return errors.New("密码少于8位")
	}
	if !MatchStr(password) {
		return errors.New("密码必须包含数字和大小写字母，可以使用特殊字符")
	}
	return nil
}

// 密码强度检测
func MatchStr(str string) bool {
	expr := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,100}$`//`^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,32}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)

	return m != nil
}