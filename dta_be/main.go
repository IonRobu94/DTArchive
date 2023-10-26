package main

import (
	"dta_be/models"
	"dta_be/services"
	"github.com/gin-gonic/gin"
)

type ServerState struct {
	Bands   []models.Band
	Albums  []models.Album
	Reviews []models.Review
}

func main() {
	bands, albums, reviews := services.GetAllData()
	state := ServerState{
		Bands:   bands,
		Albums:  albums,
		Reviews: reviews,
	}
	state.SetupRouter()
}

func (s *ServerState) SetupRouter() {
	router := gin.Default()

	router.GET("/bands", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"bands": s.Bands,
		})
	})
	router.GET("/albums", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"albums": s.Albums,
		})
	})
	router.GET("/reviews", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"reviews": s.Reviews,
		})
	})

	router.Run(":8080")
}
