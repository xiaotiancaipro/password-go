package routers

import (
	"github.com/gin-gonic/gin"
	"password-go/controllers"
)

func UserRouter(r *gin.Engine) {
	userController := controllers.UserController{}
	user := r.Group("/user")
	{
		user.POST("login", userController.Login)
		user.POST("create", userController.CreateUser)
	}
}
