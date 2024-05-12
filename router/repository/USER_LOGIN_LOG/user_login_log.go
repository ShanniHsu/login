package USER_LOGIN_LOG

import (
	"errors"
	"login/connect"
	"login/model"
)

func Create(user *model.UserLoginLog) (err error) {
	err = connect.DBConnect.Create(user).Error
	if err != nil {
		err = errors.New("UserLoginLog Create Error!")
	}
	return
}
