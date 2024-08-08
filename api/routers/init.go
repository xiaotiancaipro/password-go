package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) {
	TestRouter(r)
	UserRouter(r, db)
}
