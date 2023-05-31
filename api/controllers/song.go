package controllers

import (
	"Song_API/api/models"
	"Song_API/api/repository"
	"Song_API/api/routes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Controller struct holds a repository.SongInterface object. The controller functions use it to access the methods of the repository package
type Controller struct {
	Repo repository.SongInterface
}

// AddSong(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper AddSong(*models.Song) function to add a song in database and returns a map[string]interface{} response containing error message, status code and data
func (s *Controller) AddSong(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var song models.Song
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &song)
	if err := s.Repo.AddSong(&song); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return map[string]interface{}{
		"response": "Song added successfully",
		"data":     song,
		"status":   http.StatusOK,
	}
}

// GetAllSong(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper GetAllSong(*[]models.Song) function to get all songs from database and returns a map[string]interface{} response containing error message, status code and data
func (s *Controller) GetAllSong(ctx context.Context, req *routes.AppReq) routes.AppResp {
	fmt.Println("GetAllSong")
	var song []models.Song
	if err := s.Repo.GetAllSong(&song); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return map[string]interface{}{
		"data":   song,
		"status": http.StatusOK,
	}
}

// GetSongById(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper GetSong(*models.Song, id string) function to get a song from database and returns a map[string]interface{} response containing error message, status code and data
func (s *Controller) GetSongById(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var song models.Song
	id := req.Params["id"]
	if err := s.Repo.GetSong(&song, id); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		}
	}
	return map[string]interface{}{
		"data":   song,
		"status": http.StatusOK,
	}
}

// UpdateSong(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper UpdateSong(*models.Song) function to update a song in database and returns a map[string]interface{} response containing error message, status code and data
func (s *Controller) UpdateSong(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var song models.Song
	id := req.Params["id"]
	if err := s.Repo.GetSong(&song, id); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		}
	}
	if val, exist := req.Body["id"]; exist && val != id {
		return map[string]interface{}{
			"error":  "id mismatch; updataion not allowed",
			"status": http.StatusBadRequest,
		}
	}

	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &song)
	if err := s.Repo.UpdateSong(&song); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return map[string]interface{}{
		"response": "Song updated successfully",
		"data":     song,
		"status":   http.StatusOK,
	}
}

// DeleteSong(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper DeleteSong(*models.Song, id string) function to delete a song from database and returns a map[string]interface{} response containing error message and status code
func (s *Controller) DeleteSong(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var song models.Song
	id := req.Params["id"]
	if err := s.Repo.DeleteSong(&song, id); err != nil {
		return map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return map[string]interface{}{
		"response": fmt.Sprintf("id %s deleted", id),
		"status":   http.StatusOK,
	}
}
