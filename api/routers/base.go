package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"password-go/middleware"
	"password-go/utils"
)

var log = utils.LoggerUtil{}.Logger()
var response = utils.ResponseUtil{}

func useDB(customController func(c *gin.Context, db *gorm.DB)) func(c *gin.Context) {
	return func(c *gin.Context) {
		db, err := middleware.GetDB(c)
		if err != nil {
			log.Error(err)
			response.Error(c, fmt.Sprintf("%v", err))
			return
		}
		customController(c, db)
	}
}
