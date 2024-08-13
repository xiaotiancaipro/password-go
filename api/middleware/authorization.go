package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/types"
	"password-go/configs"
)

func getIgnoreAuthPathMap() (pathMap map[string]types.Nil) {
	paths := []string{
		"/health",
	}
	pathMap = make(map[string]types.Nil, len(paths))
	for _, path := range paths {
		pathMap[path] = types.Nil{}
	}
	return
}

func InitAuth(config configs.ConfigYaml) gin.HandlerFunc {

	return func(c *gin.Context) {

		if _, exists := getIgnoreAuthPathMap()[c.Request.URL.Path]; exists {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "No Authorization header found")
			c.Abort()
			return
		}

		var scheme, token string
		_, err := fmt.Sscanf(authHeader, "%s %s", &scheme, &token)
		if (err != nil) || (scheme != "Bearer") || (token != config.SecretKey) {
			response.Unauthorized(c, "Invalid Authorization header")
			c.Abort()
			return
		}

		c.Next()
		return

	}

}
