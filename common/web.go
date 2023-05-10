package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app Application

func init() {
	ReadConfig()
	InitDB()
}

func InitApp(port int) Application {
	if !app.Initialize {
		// gin.SetMode("release")
		app.Engine = gin.Default()
		app.Port = port
		app.Initialize = true
		app.PathPrefix = "/api"

		app.Engine.NoRoute(func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotFound, Response{
				ErrorMsg: "Not Found",
			})
		})
	}
	return app
}

func (app *Application) Run() {
	app.Engine.Run(":3000")
}
