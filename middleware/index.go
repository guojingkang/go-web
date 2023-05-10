package middleware

import (
	"learn/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(app common.Application) {
	app.Engine.Use(SessionStore())
	app.Engine.Use(resultMiddleware())
	app.Engine.Use(LoginMiddleware())
}

func SessionStore() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	return sessions.Sessions("mysession", store)
}

func resultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var response common.Response
		err, errorMsgOk := c.Get("errorMsg")

		if err != nil && errorMsgOk {
			response.Code = 500
			response.ErrorMsg = err.(string)
			c.JSON(200, response)
		} else {
			res, resultOk := c.Get("result")
			if res != nil && resultOk {
				response.Code = 200
				response.Data = res
			}
			c.JSON(200, response)
		}
	}
}

