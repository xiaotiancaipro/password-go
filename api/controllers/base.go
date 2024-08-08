package controllers

import (
	"password-go/services"
	"password-go/utils"
)

var log = utils.LoggerUtil{}.Logger()
var response = utils.ResponseUtil{}
var userService = services.UserService{}
