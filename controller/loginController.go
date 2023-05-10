package controller

import (
	"errors"
	"learn/common"
	"learn/model"
	"learn/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var loginController = common.NewRouter()

func SetupLoginController() {
	loginController.Post("/login", func(ctx *gin.Context) (any, error) {
		var input service.LoginType
		err := ctx.ShouldBind(&input)
		if err != nil || input.Account == "" || input.Password == "" {
			return nil, errors.New("入参解析失败")
		}
		user, err := service.Login(&input)
		if err != nil {
			return nil, err
		}
		session := sessions.Default(ctx)
		session.Set("user", common.JSONStringify(&user))
		session.Save()
		return user, nil
	})

	loginController.Get("/currentUser", func(ctx *gin.Context) (any, error) {
		user, exist := ctx.Get("user")
		if !exist {
			return "", errors.New("user not exist")
		}
		return user, nil
	})

	loginController.Post("/signup", func(ctx *gin.Context) (any, error) {
		var input model.User
		err := ctx.ShouldBind(&input)
		if err != nil || input.Account == "" || input.Password == "" {
			return nil, errors.New("入参解析失败")
		}
		return service.Signup(&input)
	})

	loginController.Post("/logout", func(ctx *gin.Context) (any, error) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
		return nil, nil
	})
}
