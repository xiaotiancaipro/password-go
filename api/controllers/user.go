package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"password-go/middleware"
	"password-go/models"
)

type UserController struct{}

func (u UserController) Login(c *gin.Context) {

	type input struct {
		Email    string
		Password string
	}

	type output struct {
		Exists int
		Login  int
	}

	db, err := middleware.GetDB(c)
	if err != nil {
		log.Error(err)
		response.Error(c, fmt.Sprintf("%v", err))
		return
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

	userDB := models.User{}
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

func (u UserController) CreateUser(c *gin.Context) {

	// input : models.User{}

	type output struct {
		//  1 : The user create successfully
		// -1 : Email or Password is empty
		// -2 : The email is invalid
		// -3 : This user is already exists
		// -4 : Retry, because the user create failed
		Status int
	}

	db, err := middleware.GetDB(c)
	if err != nil {
		log.Error(err)
		response.Error(c, fmt.Sprintf("%v", err))
		return
	}

	userInput := models.User{}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		message := "Invalid input"
		log.Error(message)
		response.Error(c, message)
		return
	}

	if userInput.Email == "" || userInput.Password == "" {
		message := "Email or Password is empty"
		log.Info(message)
		response.Success(c, message, output{-1})
		return
	}

	if !stringUtil.IsEmailValid(userInput.Email) {
		message := fmt.Sprintf("The email [%s] is not valid", userInput.Email)
		log.Info(message)
		response.Success(c, message, output{-2})
		return
	}

	userDB := models.User{}
	if err := db.Where("email = ?", userInput.Email).First(&userDB).Error; err == nil {
		message := fmt.Sprintf("This user [%s] is already exists", userInput.Email)
		log.Info(message)
		response.Success(c, message, output{-3})
		return
	}

	flag := userOperate.Create(db, &userInput)
	if !flag {
		message := fmt.Sprintf("This user [%s] create failed", userInput.Email)
		log.Warning(message)
		response.Success(c, message, output{-4})
		return
	}

	message := fmt.Sprintf("This user [%s] create successfully", userInput.Email)
	log.Info(message)
	response.Success(c, message, output{1})

}
