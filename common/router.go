package common

import "github.com/gin-gonic/gin"

func NewRouter() *BaseRouter {
	var router = BaseRouter{
		App: &app,
	}
	return &router
}

func (router *BaseRouter) Get(apiUrl string, handleFunc HandleFunc) {
	router.App.Engine.GET(router.App.PathPrefix+apiUrl, func(ctx *gin.Context) {
		handleRequest(handleFunc, ctx)
	})
}

func (router *BaseRouter) Post(apiUrl string, handleFunc HandleFunc) {
	router.App.Engine.POST(router.App.PathPrefix+apiUrl, func(ctx *gin.Context) {
		handleRequest(handleFunc, ctx)
	})
}

func handleRequest(handleFunc HandleFunc, ctx *gin.Context) {
	res, err := handleFunc(ctx)
	if err != nil {
		ctx.Set("errorMsg", err.Error())
		ctx.Set("result", nil)
	} else {
		ctx.Set("errorMsg", nil)
		ctx.Set("result", res)
	}

}
