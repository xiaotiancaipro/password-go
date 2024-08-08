package main

import (
	"github.com/gin-gonic/gin"
	"password-go/configs"
	"password-go/models"
	"password-go/routers"
)

func CreateApp(config configs.ConfigYaml) *gin.Engine {
	app := gin.Default()
	db := models.InitDB(config)
	routers.InitRouter(app, db)
	return app
}

func main() {
	config := configs.InitConfig()
	app := CreateApp(config)
	app.Run(":6001")
}
