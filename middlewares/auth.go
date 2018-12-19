package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tobi007/cdServer/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
			c.JSON(500, gin.H{"code": "500", "message": "Unauthorized access"})
			return
		}
		if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
			c.JSON(500, gin.H{"code": "500", "message": "Unauthorized access"})
			return
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			c.JSON(500, gin.H{"code": "401", "message": "Unauthorized access"})
			return
		}
		c.Next()
	}
}
