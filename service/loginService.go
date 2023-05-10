package service

import (
	"errors"
	"learn/model"
	"time"
)

func Signup(user *model.User) (string, error) {
	var exist = userDao.FindOne(&model.User{Account: user.Account})

	if exist.Id != 0 {
		return "", errors.New("用户账号重复")
	}

	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	
	userDao.Add(user)
	return "success", nil
}

type LoginType struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(input *LoginType) (model.User, error) {
	var user = userDao.FindOne(&model.User{Account: input.Account})
	if user.Password != input.Password {
		return model.User{}, errors.New("用户密码错误")
	}
	user.Password = ""
	return user, nil
}
