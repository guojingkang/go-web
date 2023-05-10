package controller

import (
	"errors"
	"learn/common"
	"learn/dao"
	"learn/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userController = common.NewRouter()

func StepUserController() {
	userController.Get("/user", func(ctx *gin.Context) (any, error) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			return nil, errors.New("入参解析失败")
		}
		return service.GetUser(id)
	})

	userController.Post("/user/update",func(ctx *gin.Context) (any, error) {
		input := service.UpdateUserType{}
		err := ctx.ShouldBind(&input)
		if err != nil {
			return nil, errors.New("入参解析失败")
		}
		updateErr := service.UpdateUser(&input)
		if updateErr != nil {
			return nil, updateErr
		} else {
			return nil ,nil
		}
	})

	userController.Get("/user/list", func(ctx *gin.Context) (any, error){
		var err error

		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil{
			return nil, errors.New("入参解析失败")
		}

		offset, err := strconv.Atoi(ctx.Query("offset"))
		if err != nil{
			return nil, errors.New("入参解析失败")
		}
		keyword := ctx.Query("keyword")
		return service.UserListByPage(&dao.PageQueryInput{Limit: limit,Offset: offset, Keyword: keyword}), nil
	})
}
