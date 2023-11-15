package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"artem.cz/albums"
)

func main() {

	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAllAlbums(c *gin.Context) {
	var albums = albums.getAll()
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums.getAll() {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {

	var newAlbum albums.album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	var errors string

	if newAlbum.Title == "" {
		errors += "Title can't be empty"
	}

	if newAlbum.Artist == "" {
		errors += "Artist can't be empty"
	}

	if errors != "" {
		c.IndentedJSON(http.StatusBadRequest, errors)
		return
	}

	newAlbum.ID = strconv.Itoa(len(albums.getAll()) + 1)
	// albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
