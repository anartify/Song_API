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

type Controller struct {
	Repo repository.SongInterface
}

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
