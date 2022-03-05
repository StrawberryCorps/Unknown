package main

import (
	"github.com/gin-gonic/gin"
	"unknown/strawberrycorps"
)

func main() {
	router := gin.Default()
	strawberrycorps.InitRoutes(router)
	router.Run()
}
