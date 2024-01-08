package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// DecodeJson() decodes the body of a request and puts it in the given struct
func DecodeJson[T any](r *http.Request, obj T) (*T, error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)
	return &obj, err
}

func EncodeJson[T any](obj T) string {
	ret, _ := json.MarshalIndent(obj, "", "\t")
	return string(ret)
}

func ValidateJson[T any](obj T) error {
	return validator.New().Struct(obj)
}

func GetAndValidateJsonFromRequest[T any](r *http.Request, obj T) (*T, error) {
	jsonObj, err := DecodeJson(r, obj)
	if err != nil {
		return nil, err
	}
	err = ValidateJson(jsonObj)
	return jsonObj, err
}
