package strawberrycorps

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/rest"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/api/game", rest.Games)
}
