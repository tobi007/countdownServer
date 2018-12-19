package util

import (
	"github.com/gin-gonic/gin"
)

// respondwithJSON write json response format
func RespondwithJSONAndAbort(c *gin.Context, code int,message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func RespondwithJSON(c *gin.Context, code int, message string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, data)
}