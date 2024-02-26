package common

import (
	"github.com/gin-gonic/gin"
)

func HandleFormError(key *string, message string, c *gin.Context) {
	if key != nil {
		c.JSON(400, gin.H{
			"field":   key,
			"message": message,
		})
	} else {
		c.JSON(400, gin.H{
			"message": message,
		})
	}
}
