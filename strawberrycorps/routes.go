package strawberrycorps

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/rest"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/api/game", rest.Games)
	router.POST("/api/game", rest.CreateGame)
	router.DELETE("/api/game/:id", rest.DeleteGame)
	router.GET("/api/game/:id", rest.GetGame)

	router.GET("/api/health", rest.Heatlh)
	router.GET("/api/health/:id", rest.GetGamesHealth)
}
