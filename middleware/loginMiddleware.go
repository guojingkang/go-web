package middleware

import (
	"errors"
	"learn/common"
	"learn/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkLogin(c) || checkIsWhiteUrl(c) {
			c.Next()
		} else {
			c.AbortWithError(200, errors.New("当前用户未登录"))
			c.Set("errorMsg", "当前用户未登录")
		}
	}
}

func checkLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	var userInfo model.User
	var sessionStr = session.Get("user")
	if sessionStr == nil || sessionStr == "" {
		return false
	}
	common.JSONParse(sessionStr.(string), &userInfo)
	if userInfo.Id != 0 {
		c.Set("user", &userInfo)
		return true
	} else {
		return false
	}
}

var whiteListUrl = []string{"/api/login"}

func checkIsWhiteUrl(c *gin.Context) bool {
	return common.ContainsString(whiteListUrl, c.Request.URL.Path)
}
