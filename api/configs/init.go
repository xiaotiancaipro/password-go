package configs

import "github.com/gin-gonic/gin"

func InitConfig() ConfigYaml {
	config := Config()
	gin.SetMode(gin.ReleaseMode)
	if config.Gin.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	return config
}
