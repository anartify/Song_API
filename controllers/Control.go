package controllers

import (
	"Song_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	M models.SongInterface
}

func (ctrl Controller) GetAllSong(ctx *gin.Context) {
	var song []models.Song
	err := ctrl.M.GetAllSong(&song)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, song)
	}
}

func (ctrl Controller) AddNewSong(ctx *gin.Context) {
	var song models.Song
	ctx.BindJSON(&song)
	err := ctrl.M.AddNewSong(&song)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, song)
	}
}

func (ctrl Controller) GetSong(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var song models.Song
	err := ctrl.M.GetSong(&song, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, song)
	}
}

func (ctrl Controller) UpdateSong(ctx *gin.Context) {
	var song models.Song
	id := ctx.Params.ByName("id")
	err := ctrl.M.GetSong(&song, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, song)
	}
	ctx.BindJSON(&song)
	err = ctrl.M.UpdateSong(&song, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, song)
	}
}

func (ctrl Controller) DeleteSong(ctx *gin.Context) {
	var song models.Song
	id := ctx.Params.ByName("id")
	err := ctrl.M.DeleteSong(&song, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id-" + id: "deleted"})
	}
}
