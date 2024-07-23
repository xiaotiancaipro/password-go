package routers

import (
	"github.com/gin-gonic/gin"
	"password-go/controllers"
)

func TestRouter(r *gin.Engine) {
	testController := controllers.TestController{}
	test := r.Group("/test")
	{
		test.GET("/list/:id", testController.GetTest)
		test.POST("/list/post", testController.PostTest)
	}
}
