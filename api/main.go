package main

import (
	"github.com/gin-gonic/gin"
	"password-go/configs"
	"password-go/routers"
)

func InitConfig() {
	config := configs.Config()
	gin.SetMode(gin.ReleaseMode)
	if config.Gin.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
}

func InitRouter(r *gin.Engine) {
	routers.TestRouter(r)
}

func CreateApp() *gin.Engine {
	app_ := gin.Default()
	InitRouter(app_)
	return app_
}

func main() {
	InitConfig()
	app := CreateApp()
	app.Run(":6001")
}
