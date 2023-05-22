package Routes

import (
	"Song_API/Controllers"

	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	r := gin.Default()
	grp := r.Group("/v1")
	{
		grp.GET("api", Controllers.GetAllSong)        //Get all the songs stored in database
		grp.POST("api", Controllers.AddNewSong)       //Adds a new song to SongsDB
		grp.GET("api/:id", Controllers.GetSong)       //Gets certain Song based on id
		grp.PUT("api/:id", Controllers.UpdateSong)    //Updates certain Song based on id
		grp.DELETE("api/:id", Controllers.DeleteSong) //Deletes certain Song based on id
	}
	return r
}
