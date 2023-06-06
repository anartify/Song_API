package controllers

import (
	"Song_API/pkg/controllers/utils"
	"Song_API/pkg/controllers/validation"
	"Song_API/pkg/models"
	"Song_API/pkg/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Controller struct holds a Song and Account Interface objects of repo layer. The controller functions use it to access the methods of the repository package
type Controller struct {
	SongRepo    repository.SongInterface
	AccountRepo repository.AccountInterface
}

// NewController function returns a new pointer to a Controller struct.
func NewController(songRepo repository.SongInterface, accountRepo repository.AccountInterface) *Controller {
	return &Controller{
		SongRepo:    songRepo,
		AccountRepo: accountRepo,
	}
}

// AddSong(context.Context, *utils.AppReq) function calls a helper AddSong function to add a song in database and returns a utils.AppResp response containing error message, status code and data
func (ctrl *Controller) AddSong(ctx context.Context, req *utils.AppReq) utils.AppResp {
	user := ctx.Value("user").(string)
	var song models.Song
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &song)
	if err := validation.ValidateSong(song, true); err != nil {
		return utils.AppResp{
			"error":  err,
			"status": http.StatusBadRequest,
		}
	}
	if err := ctrl.SongRepo.AddSong(&song, user); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return utils.AppResp{
		"response": "Song added successfully",
		"data":     song,
		"status":   http.StatusOK,
	}
}

// GetAllSong(context.Context, *utils.AppReq) function calls a helper GetAllSong function to get all songs from database and returns a utils.AppResp response containing error message, status code and data
func (ctrl *Controller) GetAllSong(ctx context.Context, req *utils.AppReq) utils.AppResp {
	user := ctx.Value("user").(string)
	var song []models.Song
	if err := ctrl.SongRepo.GetAllSong(&song, user); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return utils.AppResp{
		"data":   song,
		"status": http.StatusOK,
	}
}

// GetSongById(context.Context, *utils.AppReq) function calls a helper GetSong function to get a song from database and returns a utils.AppResp response containing error message, status code and data
func (ctrl *Controller) GetSongById(ctx context.Context, req *utils.AppReq) utils.AppResp {
	user := ctx.Value("user").(string)
	var song models.Song
	id := req.Params["id"]
	if err := ctrl.SongRepo.GetSong(&song, id, user); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		}
	}
	return utils.AppResp{
		"data":   song,
		"status": http.StatusOK,
	}
}

// UpdateSong(context.Context, *utils.AppReq) function calls a helper UpdateSong to update a song in database and returns a utils.AppResp response containing error message, status code and data
func (ctrl *Controller) UpdateSong(ctx context.Context, req *utils.AppReq) utils.AppResp {
	user := ctx.Value("user").(string)
	var song models.Song
	id := req.Params["id"]
	if err := ctrl.SongRepo.GetSong(&song, id, user); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		}
	}
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &song)
	if err := validation.ValidateSong(song, false); err != nil {
		return utils.AppResp{
			"error":  err,
			"status": http.StatusBadRequest,
		}
	}
	if val, exist := req.Body["id"]; exist && fmt.Sprintf("%v", val) != id {
		return utils.AppResp{
			"error":  "id mismatch, updation not allowed",
			"status": http.StatusBadRequest,
		}
	}
	if val, exist := req.Body["user"]; exist && user != val {
		return utils.AppResp{
			"error":  "user updation not allowed",
			"status": http.StatusBadRequest,
		}
	}
	if err := ctrl.SongRepo.UpdateSong(&song); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return utils.AppResp{
		"response": "Song updated successfully",
		"data":     song,
		"status":   http.StatusOK,
	}
}

// DeleteSong(context.Context, *utils.AppReq) function calls a helper DeleteSong function to delete a song from database and returns a utils.AppResp response containing error message and status code
func (ctrl *Controller) DeleteSong(ctx context.Context, req *utils.AppReq) utils.AppResp {
	user := ctx.Value("user").(string)
	var song models.Song
	id := req.Params["id"]
	if err := ctrl.SongRepo.DeleteSong(&song, id, user); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		}
	}
	return utils.AppResp{
		"response": "id " + id + " deleted by " + user,
		"status":   http.StatusOK,
	}
}
