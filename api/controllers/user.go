package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"password-go/models"
)

type UserController struct{}

func (u UserController) CheckUser(c *gin.Context, db *gorm.DB) {

	type input struct {
		Email    string
		Password string
	}

	type output struct {
		Exists int
		Login  int
	}

	if userOperate.IsTableEmpty(db) {
		message := "The table [user] is empty"
		log.Info(message)
		response.Success(c, message, output{0, 0})
		return
	}

	userInput := input{}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		message := "Invalid input"
		log.Error(message)
		response.Error(c, message)
		return
	}

	var userDB models.User
	if err := db.Where("email = ?", userInput.Email).First(&userDB).Error; err != nil {
		message := fmt.Sprintf("This user [%s] is not exists", userInput.Email)
		log.Info(message)
		response.Success(c, message, output{0, 0})
		return
	}

	if userInput.Password != userDB.Password {
		message := fmt.Sprintf("The user [%s] entered an incorrect password", userInput.Email)
		log.Info(message)
		response.Success(c, message, output{1, 0})
		return
	}

	message := "Login successful"
	log.Info(message)
	response.Success(c, message, output{1, 1})

}
