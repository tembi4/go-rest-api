package main

import (
	"net/http"
	"strconv"

	"artem.cz/albums"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.GetAll())
}

func getAlbumById(c *gin.Context) {

	idParamAsString := c.Param("id")
	id, err := strconv.Atoi(idParamAsString)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idParamAsString + " is not valid number"})
		return
	}

	album, err := albums.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {

	var newAlbum albums.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not valid Album"})
		return
	}

	var album, errors = albums.CreateAlbum(newAlbum)

	if errors != "" {
		c.IndentedJSON(http.StatusBadRequest, errors)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
