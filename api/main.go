package main

import (
	"github.com/gin-gonic/gin"
	"password-go/configs"
)

func InitConfig() {
	config := configs.Config()
	gin.SetMode(gin.ReleaseMode)
	if config.Gin.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
}

func CreateApp() *gin.Engine {
	app_ := gin.Default()
	configs.InitDB()
	configs.InitRouter(app_)
	return app_
}

func main() {
	InitConfig()
	app := CreateApp()
	app.Run(":6001")
}
