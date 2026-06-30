package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

var albums = []album{
	{"1", "BRATLAND", "MACAN", 10000},
	{"2", "Vremy kolokol'chikov", "BOOKER", 1},
	{"3", "G-Woman", "ICEGERGERT", 20000},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var newAlbum album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, v := range albums {
		if v.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, v := range albums {
		if v.ID == id {
			albums = slices.Delete(albums, i, i+1)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", updateAlbumByID)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}
