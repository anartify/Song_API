package controllers

import (
	"Song_API/api/models"
	"Song_API/api/routes"
	"context"
	"encoding/json"
	"net/http"
)

// CreateAccount(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper CreateAccount function to create an account in database and returns a routes.AppResp response containing error message, status code and account data
func (ctrl *Controller) CreateAccount(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := ctrl.AccountRepo.CreateAccount(&acc); err != nil {
		return routes.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return routes.AppResp{
		"response": "Account created successfully",
		"account":  acc,
		"status":   http.StatusOK,
	}
}

// GetAccount(context.Context, *routes.AppReq) is a gin.HandlerFunc that calls a helper GetAccount function to get an account from database and returns a routes.AppResp response containing error message, status code, account data and authentication token
func (ctrl *Controller) GetAccount(ctx context.Context, req *routes.AppReq) routes.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	token, err := ctrl.AccountRepo.GetAccount(&acc)
	if err != nil {
		return routes.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return routes.AppResp{
		"account": acc,
		"token":   token,
		"status":  http.StatusOK,
	}
}
