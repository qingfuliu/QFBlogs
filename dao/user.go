package dao

import (
	"errors"
	"goProject/models"
)

func InsertUser(user *models.User) error {
	if err := dbConn.Create(user).Error; err != nil {
		return ErrorInsertFailed
	}
	return nil
}

func IfUsersExisted(username string) (bool, error) {
	return IfIsExisted("username", username, models.User{}.TableName())
}

func IfCertified(user *models.UserLoginForm) (err error) {
	var md5PassWord string
	if err = dbConn.Model(user).Select("password").Where("username=?", user.UserName).Find(&md5PassWord).Error; err != nil {
		return
	}
	if md5PassWord != user.PassWord {
		err = errors.New("password not match")
	}
	return
}
