package common

import "github.com/gin-gonic/gin"

type Application struct {
	Engine     *gin.Engine
	Port       int
	Initialize bool
	PathPrefix string
}

type BaseRouter struct {
	App *Application
}

type Response struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"errorMsg"`
	Data     any    `json:"data"`
}

type HandleFunc func(ctx *gin.Context) (any, error)
