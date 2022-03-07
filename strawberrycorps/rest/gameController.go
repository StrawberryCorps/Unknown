package rest

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/service"
	"unknown/strawberrycorps/ustruct"
)

func Games(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	service.GetGamesService()
}

func CreateGame(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	service.CreateGameService()
}

func DeleteGame(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var gameUri ustruct.GameUri
	if err := c.ShouldBindUri(&gameUri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

}

func GetGame(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var gameUri ustruct.GameUri
	if err := c.ShouldBindUri(&gameUri); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

}
