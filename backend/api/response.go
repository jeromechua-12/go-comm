package api

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data  any       `json:"data"`
	Error *APIError `json:"error"`
}

type APIError struct {
	ErrorType ErrorType         `json:"errorCode"`
	Message   string            `json:"message"`
	Details   map[string]string `json:"details"`
}

func WriteSuccess(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIResponse{
		Data:  data,
		Error: nil,
	})
}

func WriteError(w http.ResponseWriter, status int, errType ErrorType, message string, details map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := APIError{
		ErrorType: errType,
		Message:   message,
		Details:   details,
	}

	json.NewEncoder(w).Encode(APIResponse{
		Data:  nil,
		Error: &err,
	})
}
