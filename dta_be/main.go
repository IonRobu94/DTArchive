package main

import (
	"dta_be/models"
	"dta_be/services"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
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
func toInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("toInterfaceSlice() given a non-slice type")
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func paginate(c *gin.Context, fullData []interface{}) []interface{} {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	start := (page - 1) * limit
	end := start + limit

	if start > len(fullData) {
		return []interface{}{} // ritorna un array vuoto se l'offset supera la lunghezza dei dati
	}

	if end > len(fullData) {
		end = len(fullData) // Adegua l'endpoint se supera la lunghezza dei dati
	}

	return fullData[start:end]
}
func (s *ServerState) SetupRouter() {
	router := gin.Default()

	router.GET("/bands", func(c *gin.Context) {
		paginatedData := paginate(c, toInterfaceSlice(s.Bands)) // Cast bands a []interface{}
		c.JSON(200, gin.H{
			"bands": paginatedData,
		})
	})
	router.GET("/albums", func(context *gin.Context) {
		paginatedData := paginate(context, toInterfaceSlice(s.Albums))
		context.JSON(200, gin.H{
			"albums": paginatedData,
		})
	})
	router.GET("/reviews", func(context *gin.Context) {
		paginatedData := paginate(context, toInterfaceSlice(s.Reviews))
		context.JSON(200, gin.H{
			"reviews": paginatedData,
		})
	})

	router.Run(":8080")
}
