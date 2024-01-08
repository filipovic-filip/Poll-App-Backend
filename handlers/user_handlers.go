package handlers

import (
	"fmt"
	"net/http"

	"filip.filipovic/polling-app/dto"
	"filip.filipovic/polling-app/logging"
	"filip.filipovic/polling-app/services"
	"filip.filipovic/polling-app/utils"
	"github.com/julienschmidt/httprouter"
)

// RegisterUser() Used to register a user and put his information in the database
//
// POST Request - Gets User as a body
func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userReq, err := utils.GetAndValidateJsonFromRequest(r, dto.UserRequest{})
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := services.RegisterUser(r.Context(), userReq)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	logging.Info(fmt.Sprintf("User with username '%s' registered", user.Username))
	fmt.Fprint(w, utils.EncodeJson(user))
}

// LoginUser() Used to log a user in
//
// POST Request
func LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginReq, err := utils.GetAndValidateJsonFromRequest(r, dto.LoginRequest{})
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := services.LoginUser(r.Context(), loginReq)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	logging.Info(fmt.Sprintf("User with username '%s' logged in", user.Username))
	fmt.Fprint(w, utils.EncodeJson(user))
}
