package controllers

import (
	"Song_API/pkg/controllers/utils"
	"Song_API/pkg/controllers/validation"
	"Song_API/pkg/models"
	"context"
	"encoding/json"
	"net/http"
)

// CreateAccount(context.Context, *routes.AppReq) function calls a helper CreateAccount function to create an account in database and returns a utils.AppResp response containing error message, status code and account data
func (ctrl *Controller) CreateAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := validation.ValidateAccount(acc); err != nil {
		return utils.AppResp{
			"error":  err,
			"status": http.StatusBadRequest,
		}
	}
	if err := ctrl.AccountRepo.CreateAccount(&acc); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return utils.AppResp{
		"response": "Account created successfully",
		"account":  acc,
		"status":   http.StatusOK,
	}
}

// GetAccount(context.Context, *routes.AppReq) function calls a helper GetAccount function to get an account from database and returns a utils.AppResp response containing error message, status code, account data and authentication token
func (ctrl *Controller) GetAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := validation.ValidateAccount(acc); err != nil {
		return utils.AppResp{
			"error":  err,
			"status": http.StatusBadRequest,
		}
	}
	token, err := ctrl.AccountRepo.GetAccount(&acc)
	if err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	return utils.AppResp{
		"account": acc,
		"token":   token,
		"status":  http.StatusOK,
	}
}
