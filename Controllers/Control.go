package Controllers

import (
	"Song_API/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSong(c *gin.Context) {
	var song []Models.Song
	err := Models.GetAllSong(&song)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func AddNewSong(c *gin.Context) {
	var song Models.Song
	c.BindJSON(&song)
	err := Models.AddNewSong(&song)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func GetSong(c *gin.Context) {
	id := c.Params.ByName("id")
	var song Models.Song
	err := Models.GetSong(&song, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func UpdateSong(c *gin.Context) {
	var song Models.Song
	id := c.Params.ByName("id")
	err := Models.GetSong(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, song)
	}
	c.BindJSON(&song)
	err = Models.UpdateSong(&song, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func DeleteSong(c *gin.Context) {
	var song Models.Song
	id := c.Params.ByName("id")
	err := Models.DeleteSong(&song, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
