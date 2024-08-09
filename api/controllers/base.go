package controllers

import (
	"password-go/models"
	"password-go/utils"
)

var userOperate = models.UserOperate{}

var log = utils.LoggerUtil{}.Logger()
var response = utils.ResponseUtil{}
var stringUtil = utils.StringUtil{}
