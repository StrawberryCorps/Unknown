package rest

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/ustruct"
)

func Games(c *gin.Context) {

}

func CreateGame(c *gin.Context) {

}

func DeleteGame(c *gin.Context) {
	var gameUri ustruct.GameUri
	if err := c.ShouldBindUri(&gameUri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

}

func GetGame(c *gin.Context) {
	var gameUri ustruct.GameUri
	if err := c.ShouldBindUri(&gameUri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

}
