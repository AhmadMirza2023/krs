package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AhmadMirza2023/krs/types"
	"github.com/go-playground/validator"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, statucCode int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statucCode)
	return json.NewEncoder(w).Encode(value)
}

func FormatResponse(w http.ResponseWriter, statusCode int, status string, message string, data any) {
	response := types.FullResponse{
		Meta: types.MetaData{
			Code:    statusCode,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
	WriteJson(w, statusCode, response)
}
