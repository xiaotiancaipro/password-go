package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUtil struct{}

func (r ResponseUtil) Build(c *gin.Context, code int, message string, data any) {
	c.JSON(code, gin.H{"code": code, "message": message, "data": data})
}

func (r ResponseUtil) Unauthorized(c *gin.Context, message string) {
	log.Error(message)
	r.Build(c, http.StatusUnauthorized, "Error: "+message, nil)
}

func (r ResponseUtil) Success(c *gin.Context, message string, data any) {
	r.Build(c, http.StatusOK, "Success: "+message, data) // 200
}

func (r ResponseUtil) Error(c *gin.Context, message string) {
	r.Build(c, http.StatusBadRequest, "Error: "+message, nil) // 400
}
