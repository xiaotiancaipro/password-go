package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUtil struct{}

func (responseUtil ResponseUtil) Build(c *gin.Context, code int, message string, data any) {
	c.JSON(code, gin.H{"code": code, "message": message, "data": data})
}

func (responseUtil ResponseUtil) Success(c *gin.Context, message string, data any) {
	responseUtil.Build(c, http.StatusOK, "Success: "+message, data) // 200
}

func (responseUtil ResponseUtil) Error(c *gin.Context, message string) {
	responseUtil.Build(c, http.StatusBadRequest, "Error: "+message, nil) // 400
}
