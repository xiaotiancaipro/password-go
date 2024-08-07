package configs

import (
	"github.com/gin-gonic/gin"
	"password-go/routers"
)

func InitRouter(r *gin.Engine) {
	routers.TestRouter(r)
}
