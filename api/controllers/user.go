package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"password-go/models"
)

type UserController struct{}

func (u UserController) Login(c *gin.Context, db *gorm.DB) {

	var input struct {
		Email    string
		Password string
	}

	/**
	[Exists: 0 & Login: 0]: User is not exists
	[Exists: 1 & Login: 0]: The user password is incorrect
	[Exists: 1 & Login: 1]: Login successful
	*/
	type output struct {
		Exists int
		Login  int
	}

	if userOperate.IsTableEmpty(db) {
		response.Success(c, "The table [user] is empty", output{0, 0})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, "Invalid input")
		return
	}

	userDB := models.User{}
	if err := db.Where("email = ?", input.Email).First(&userDB).Error; err != nil {
		response.Success(c, fmt.Sprintf("This user [%s] is not exists", input.Email), output{0, 0})
		return
	}

	if input.Password != userDB.Password {
		response.Success(c, fmt.Sprintf("The user [%s] entered an incorrect password", input.Email), output{1, 0})
		return
	}

	response.Success(c, "Login successful", output{1, 1})

}

func (u UserController) CreateUser(c *gin.Context, db *gorm.DB) {

	var input = models.User{}

	/**
	 1 : The user create successfully
	-1 : Email or Password is empty
	-2 : The email is invalid
	-3 : This user is already exists
	-4 : Retry, because the user create failed
	*/
	type output struct {
		Status int
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, "Invalid input")
		return
	}

	if (input.Email == "") || (input.Password == "") {
		response.Success(c, "Email or Password is empty", output{-1})
		return
	}

	if !stringUtil.IsEmailValid(input.Email) {
		response.Success(c, fmt.Sprintf("The email [%s] is not valid", input.Email), output{-2})
		return
	}

	userDB := models.User{}
	if err := db.Where("email = ?", input.Email).First(&userDB).Error; err == nil {
		response.Success(c, fmt.Sprintf("This user [%s] is already exists", input.Email), output{-3})
		return
	}

	flag := userOperate.Create(db, &input)
	if !flag {
		response.Warning(c, fmt.Sprintf("This user [%s] create failed", input.Email), output{-4})
		return
	}

	response.Success(c, fmt.Sprintf("This user [%s] create successfully", input.Email), output{1})

}
