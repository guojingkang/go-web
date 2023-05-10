package main

import (
	"learn/common"
	"learn/controller"
	"learn/middleware"
)

func main() {
	app := common.InitApp(3000)
	middleware.SetupMiddleware(app)
	controller.StepController()
	app.Run()
}
