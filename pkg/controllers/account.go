package controllers

import (
	"Song_API/pkg/controllers/utils"
	"Song_API/pkg/controllers/validation"
	"Song_API/pkg/models"
	appUtility "Song_API/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
)

// CreateAccount(context.Context, *utils.AppReq) function calls a helper CreateAccount function to create an account in database and returns a utils.AppResp response containing error message, status code and account data
func (ctrl *Controller) CreateAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := validation.ValidateCreateAccount(acc); err != nil {
		return utils.AppResp{
			"error":  err,
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

// GetAccount(context.Context, *utils.AppReq) function calls a helper GetAccount function to get an account from database and returns a utils.AppResp response containing error message, status code, account data and authentication token
func (ctrl *Controller) GetAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := validation.ValidateGetAccount(acc); err != nil {
		return utils.AppResp{
			"error":  err,
			"status": http.StatusBadRequest,
		}
	}
	if err := ctrl.AccountRepo.GetAccount(&acc); err != nil {
		return utils.AppResp{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		}
	}
	token, remTime, _ := appUtility.GenerateToken(&acc)
	ctrl.AccountCache.Set("token", token, remTime)
	return utils.AppResp{
		"account": acc,
		"token":   token,
		"status":  http.StatusOK,
	}
}

// GetAllAccount(context.Context, *utils.AppReq) function calls a helper GetAllAccount function to get user and role field of all accounts from database and returns a utils.AppResp response containing error message, status code and accounts data
func (ctrl *Controller) GetAllAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var resAcc []map[string]interface{}
	err := ctrl.AccountCache.Get("all", &resAcc)
	if err != nil {
		var acc []models.Account
		if err := ctrl.AccountRepo.GetAllAccount(&acc); err != nil {
			return utils.AppResp{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			}
		}
		for _, account := range acc {
			resAcc = append(resAcc, map[string]interface{}{
				"user": account.GetUser(),
				"role": account.GetRole()})
		}
		val, _ := json.Marshal(resAcc)
		ctrl.AccountCache.Set("all", string(val))
	}
	return utils.AppResp{
		"accounts": resAcc,
		"status":   http.StatusOK,
	}
}

// UpdateRole(context.Context, *utils.AppReq) function calls a helper UpdateRole function to update role of an account in database and returns a utils.AppResp response containing error, response message, status code and account data
func (ctrl *Controller) UpdateRole(ctx context.Context, req *utils.AppReq) utils.AppResp {
	var acc models.Account
	bodyBytes, _ := json.Marshal(req.Body)
	json.Unmarshal(bodyBytes, &acc)
	if err := validation.ValidateUpdateRole(acc); err != nil {
		return utils.AppResp{
			"error":  err,
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
	ctrl.AccountCache.Delete("all")
	return utils.AppResp{
		"response": "Account role updated successfully",
		"account":  resAcc,
		"status":   http.StatusOK,
	}
}

// DeleteAccount(context.Context, *utils.AppReq) function allows account owners to delete their account. It also allows admin user to delete any users account. It calls a helper DeleteAccount function to delete an account from database and returns a utils.AppResp response containing error message, status code and response message
func (ctrl *Controller) DeleteAccount(ctx context.Context, req *utils.AppReq) utils.AppResp {
	tokenClaims := ctx.Value("token").(map[string]interface{})
	user := tokenClaims["user"].(string)
	role := tokenClaims["role"].(string)
	var acc models.Account
	acc.SetUser(req.Query["user"])
	if role == "admin" || user == acc.GetUser() {
		if err := ctrl.AccountRepo.DeleteAccount(&acc); err != nil {
			return utils.AppResp{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			}
		}
		ctrl.AccountCache.Delete("all")
		if user == acc.GetUser() {
			ctrl.AccountCache.Delete("token")
		}
		return utils.AppResp{
			"response": "Account associated with " + acc.GetUser() + " deleted successfully by " + user,
			"status":   http.StatusOK,
		}
	}
	return utils.AppResp{
		"error":  "Only admin or account owner can delete this account",
		"status": http.StatusUnauthorized,
	}
}
