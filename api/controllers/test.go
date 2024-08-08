package controllers

import (
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (t TestController) GetTest(c *gin.Context) {
	id := c.Param("id")
	log.Info("GetTest" + id)
	response.Success(c, "", gin.H{"id": id})
}

func (t TestController) PostTest(c *gin.Context) {

	//id := c.PostForm("id")
	params := make(map[string]any)
	err := c.BindJSON(&params)
	if err != nil {
		response.Error(c, "")
		return
	}

	log.Info("Success: PostTest")

	response.Success(c, "", params)

}
