package main

import (
	"github.com/gin-gonic/gin"
	"password-go/configs"
	"password-go/middleware"
	"password-go/routers"
)

func CreateApp(config configs.ConfigYaml) *gin.Engine {
	app := gin.Default()
	app.Use(middleware.InitAuth(config))
	app.Use(middleware.InitDB(config))
	routers.InitRouter(app)
	return app
}

func main() {
	config := configs.InitConfig()
	app := CreateApp(config)
	app.Run(":6001")
}
