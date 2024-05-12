package USER

import (
	"errors"
	"login/connect"
	"login/model"
)

func Create(user *model.User) (err error) {
	err = connect.DBConnect.Create(user).Error
	if err != nil {
		err = errors.New("User Create Error!")
	}
	return
}

func GetByID(ID int64) (user *model.User, err error) {
	user = new(model.User)
	err = connect.DBConnect.Where("id", ID).First(user).Error
	return
}

func GetByAccount(account string) (user *model.User, err error) {
	user = new(model.User)
	err = connect.DBConnect.Where("account", account).First(user).Error
	return
}

func GetByEmail(email string) (user *model.User, err error) {
	user = new(model.User)
	err = connect.DBConnect.Where("email", email).First(user).Error
	return
}

func GetByToken(token string) (user *model.User, err error) {
	user = new(model.User)
	err = connect.DBConnect.Where("token", token).First(user).Error
	return
}

func GetByTempToken(tempToken string) (user *model.User, err error) {
	user = new(model.User)
	err = connect.DBConnect.Where("temp_token", tempToken).First(user).Error
	return
}

func Update(user *model.User, newdata map[string]interface{}) (err error) {
	err = connect.DBConnect.Model(user).Updates(newdata).Error
	if err != nil {
		err = errors.New("User Update Error!")
	}
	return
}
