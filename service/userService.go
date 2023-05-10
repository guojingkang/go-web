package service

import (
	"learn/dao"
	"learn/model"
)

var userDao = dao.UseDao()

func Test() int {
	return 1234
}
func GetUser(id int) (model.User, error) {
	return userDao.GetById(id)
}

type UpdateUserType struct {
	Id     int              `json:"id,omitempty"`
	Name   string           `json:"name,omitempty"`
	Phone  string           `json:"phone,omitempty"`
	Detail model.UserDetail `json:"detail,omitempty"`
}

func UpdateUser(input *UpdateUserType) error {
	userDao.Update(input.Id, model.User{
		Name:         input.Name,
		Phone:        input.Phone,
		DetailStruct: input.Detail,
	})
	return nil
}

func UserListByPage(input *dao.PageQueryInput) dao.UserPageResult {
	return userDao.GetListByPage(input, &model.User{})
}
