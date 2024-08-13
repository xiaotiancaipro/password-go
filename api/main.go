package main

import (
	"github.com/gin-gonic/gin"
	"password-go/configs"
	"password-go/middleware"
	"password-go/routers"
)

func InitConfig() configs.ConfigYaml {
	config := configs.Config()
	gin.SetMode(gin.ReleaseMode)
	if config.Gin.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	return config
}

func InitRouter(r *gin.Engine) {
	routers.UserRouter(r)
}

func CreateApp() *gin.Engine {
	app := gin.Default()
	config := InitConfig()
	app.Use(middleware.InitAuth(config))
	app.Use(middleware.InitDB(config))
	InitRouter(app)
	return app
}

func main() {
	app := CreateApp()
	app.Run(":6001")
}
