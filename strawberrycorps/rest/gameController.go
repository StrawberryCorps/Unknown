package rest

import "github.com/gin-gonic/gin"

func Games(c *gin.Context) {
	c.JSON(200, gin.H{
		"games": "NULL",
	})
}
