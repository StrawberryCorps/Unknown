package strawberrycorps

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps/rest"
)

/*
URLs:
- /api/game : GET liste des serveurs de jeu en cours de fonctionnement
- /api/game : POST pour lancer un serveur de jeu
- /api/game/:id : DELETE pour supprimer un serveur de jeu
- /api/game/:id : GET pour récupérer un serveur de jeu

- /api/health : GET infos sur tous les serveurs de jeu
- /api/health/:id : GET infos sur un serveur de jeu
*/

func InitRoutes(router *gin.Engine) {
	router.GET("/api/game", rest.Games)
	router.POST("/api/game", rest.CreateGame)
	router.DELETE("/api/game/:id", rest.DeleteGame)
	router.GET("/api/game/:id", rest.GetGame)

	router.GET("/api/health", rest.Heatlh)
	router.GET("/api/health/:id", rest.GetGamesHealth)
}
