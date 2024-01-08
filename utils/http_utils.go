package utils

import (
	"errors"
	"net/http"
	"strconv"

	"filip.filipovic/polling-app/logging"
	"github.com/julienschmidt/httprouter"
)

func SendErrorResponse(w http.ResponseWriter, err error, errCode int) {
	logging.Err(err)
	http.Error(w, err.Error(), errCode)
}

func ExtractIdFromParams(w http.ResponseWriter, p httprouter.Params) (int, error) {
	id := p.ByName("id")
	if len(id) == 0 {
		return 0, errors.New("ID Missing")
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return idNum, nil
}