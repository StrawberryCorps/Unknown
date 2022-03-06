package rest

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/ustruct"
)

func Heatlh(c *gin.Context) {

}

func GetGamesHealth(c *gin.Context) {
	var gameUri ustruct.GameUri
	if err := c.ShouldBindUri(&gameUri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

}
