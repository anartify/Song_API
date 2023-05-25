package Routes

import (
	"Song_API/Controllers"
	"Song_API/Models"

	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	ctrl := Controllers.Controller{M: Models.SongModel{}}
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
