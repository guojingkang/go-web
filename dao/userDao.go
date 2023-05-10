package dao

import (
	"errors"
	"fmt"
	"learn/common"
	"learn/model"
	"sync"

	"gorm.io/gorm"
)

type UserDTO = Dao

func UseDao() UserDTO {
	var userDao = UserDTO{}
	userDao.Db = common.GetDB()
	return userDao
}

func (userDao UserDTO) GetById(id int) (model.User, error) {
	var user model.User
	err := userDao.Db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("找不到对应的记录")
		} else {
			return user, errors.New("查询异常")
		}
	} else {
		normalizeStr2Detail(&user)
		fmt.Println(user)
		return user, nil
	}
}

func (userDao UserDTO) Add(user *model.User) {
	userDao.Db.Create(user)
}

func (userDao UserDTO) FindOne(user *model.User) model.User {
	var res model.User
	userDao.Db.Where(user).First(&res)
	return res
}

func (userDao UserDTO) GetList(where *model.User, tx *gorm.DB) []model.User {
	var list []model.User
	userDao.Db.Where(where).Find(&list)
	return list
}

type UserPageResult struct {
	Total int64        `json:"total"`
	List  []model.User `json:"list"`
}

func (userDao UserDTO) GetListByPage(input *PageQueryInput, where *model.User) UserPageResult {
	var result UserPageResult
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		var count int64
		userDao.Db.Model(&model.User{}).Count(&count)
		result.Total = count
	}()

	go func() {
		defer wg.Done()
		var users []model.User
		userDao.Db.Limit(input.Limit).Offset(input.Offset).Where(where).Find(&users)
		result.List = users
	}()
	wg.Wait()
	return result
}

func (userDao UserDTO) Update(id int, user model.User) {
	normalizeDetail2Str(&user)
	userDao.Db.Model(&model.User{Id: id}).Updates(user)
}

func normalizeDetail2Str(user *model.User) {
	user.Detail = common.JSONStringify(user.DetailStruct)
}

func normalizeStr2Detail(user *model.User) {
	var detailStruct = model.UserDetail{}
	common.JSONParse(user.Detail, &detailStruct)
	user.DetailStruct = detailStruct
	user.Detail = ""
}
