package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"password-go/configs"
	"password-go/extensions"
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

func InitRouter(r *gin.Engine, db *gorm.DB) {
	routers.UserRouter(r, db)
}

func CreateApp(config configs.ConfigYaml) *gin.Engine {
	app := gin.Default()
	db := extensions.InitDB(config)
	InitRouter(app, db)
	return app
}

func main() {
	config := InitConfig()
	app := CreateApp(config)
	app.Run(":6001")
}
