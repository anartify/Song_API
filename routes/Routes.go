package routes

import (
	"Song_API/controllers"
	"Song_API/models"

	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	ctrl := controllers.Controller{M: models.SongModel{}}
	r := gin.Default()
	grp := r.Group("/v1")
	{
		grp.GET("api", ctrl.GetAllSong)        //Get all the songs stored in database
		grp.POST("api", ctrl.AddNewSong)       //Adds a new song to SongsDB
		grp.GET("api/:id", ctrl.GetSong)       //Gets certain Song based on id
		grp.PUT("api/:id", ctrl.UpdateSong)    //Updates certain Song based on id
		grp.DELETE("api/:id", ctrl.DeleteSong) //Deletes certain Song based on id
	}
	return r
}
