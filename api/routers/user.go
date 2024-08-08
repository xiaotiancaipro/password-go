package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"password-go/controllers"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {
	userController := controllers.UserController{}
	user := r.Group("/user")
	{
		user.POST("check", func(c *gin.Context) { userController.CheckUser(c, db) })
	}
}
