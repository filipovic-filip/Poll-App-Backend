package handlers

import (
	"fmt"
	"net/http"

	"filip.filipovic/polling-app/dto"
	"filip.filipovic/polling-app/services"
	"filip.filipovic/polling-app/utils"
	"github.com/julienschmidt/httprouter"
)

// GetUsersForPollOption() Returns all users that voted for that poll option
//
// GET Request
func GetUsersForPollOption(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := utils.ExtractIdFromParams(w, p)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	users, err := services.GetUsersForPollOption(r.Context(), id)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, utils.EncodeJson(users))
}

// VoteForPollOption() Used to vote for a poll option. If the user already voted on this poll, error is returned
//
// POST Request
func VoteForPollOption(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	voteRequest, err := utils.GetAndValidateJsonFromRequest(r, dto.VoteForPollOptionRequest{})
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	pollOption, err := services.VoteForPollOption(r.Context(), voteRequest)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, utils.EncodeJson(pollOption))
}
