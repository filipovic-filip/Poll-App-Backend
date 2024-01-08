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

// GetPoll() Returns a poll based on it's id
//
// GET Request
func GetPoll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := utils.ExtractIdFromParams(w, p)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	username := r.URL.Query().Get("username")
	pollResp, err := services.GetPollByID(r.Context(), id, username)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, utils.EncodeJson(pollResp))
}

// GetPolls() Returns all polls.
//
// Query params:
// name - returns all polls whose names start with the given string
//
// GET request
func GetPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	namePrefix := r.URL.Query().Get("name")
	polls, err := services.GetPolls(r.Context(), namePrefix)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, utils.EncodeJson(polls))
}

// GetPollsCreatedByUser() Returns all polls that were created by the given user based on his userId
//
// GET Request
func GetPollsCreatedByUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createdByUsername := p.ByName("createdByUsername")

	polls, err := services.GetPollsByCreatedByUsername(r.Context(), createdByUsername)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, utils.EncodeJson(polls))
}

// CreatePoll() Created a poll
//
// POST Request
func CreatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pollReq, err := utils.GetAndValidateJsonFromRequest(r, dto.PollRequest{})
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	poll, err := services.CreatePoll(r.Context(), pollReq)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	logging.Info(fmt.Sprintf("Poll named '%s' created", poll.Name))
	fmt.Fprint(w, utils.EncodeJson(poll))
}

// ModifyPoll() Modifies a given poll. Can change the name and description of a poll
// as well as add or remove poll options
//
// PUT Request
func ModifyPoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	modifyPollReq, err := utils.GetAndValidateJsonFromRequest(r, dto.ModifyPollRequest{})
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	poll, err := services.ModifyPoll(r.Context(), modifyPollReq)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	logging.Info(fmt.Sprintf("Poll named '%s' modified", poll.Name))
	fmt.Fprint(w, utils.EncodeJson(poll))
}
