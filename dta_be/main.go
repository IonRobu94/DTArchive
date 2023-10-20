package main

import (
	"dta_be/models"
	"dta_be/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

var bands []models.Band

func main() {
	fmt.Print("Hello World")
	bands = services.LoadBandsFomCSV()
	setupRouter()
}

func setupRouter() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/bands", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"bands": bands,
		})
	})

	router.Run(":8080")
}
