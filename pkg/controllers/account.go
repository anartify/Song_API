package controllers

import (
	"Song_API/pkg/apperror"
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
	var customErr apperror.CustomError
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	errUser := validation.ValidateUser(acc)
	errPass := validation.ValidatePassword(acc)
	if errUser != nil || errPass != nil {
		return utils.AppResp{
			"error":  customErr.Combine([]error{errUser, errPass}).Error(),
			"status": http.StatusBadRequest,
		}
	}
	acc.Role = "general"
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
	var customErr apperror.CustomError
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	errUser := validation.ValidateUser(acc)
	errPass := validation.ValidatePassword(acc)
	if errUser != nil || errPass != nil {
		return utils.AppResp{
			"error":  customErr.Combine([]error{errUser, errPass}).Error(),
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

func (ctrl *Controller) GetAllAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc []models.Account
	if err := ctrl.AccountRepo.GetAllAccount(&acc); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	var resAcc []map[string]interface{}
	for _, account := range acc {
		resAcc = append(resAcc, map[string]interface{}{
			"user": account.GetUser(),
			"role": account.GetRole()})
	}
	return utils.AppResp{
		"accounts": resAcc,
		"status":   http.StatusOK,
	}
}

func (ctrl *Controller) UpdateRole(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	var customErr apperror.CustomError
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	errUser := validation.ValidateUser(acc)
	errRole := validation.ValidateRole(acc)
	if errUser != nil || errRole != nil {
		return utils.AppResp{
			"error":  customErr.Combine([]error{errUser, errRole}).Error(),
			"status": http.StatusBadRequest,
		}
	}
	if err := ctrl.AccountRepo.UpdateRole(&acc); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	resAcc := map[string]interface{}{
		"user": acc.GetUser(),
		"role": acc.GetRole(),
	}
	return utils.AppResp{
		"response": "Account role updated successfully",
		"account":  resAcc,
		"status":   http.StatusOK,
	}
}
