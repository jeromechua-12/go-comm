package internal

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data  any       `json:"data"`
	Error *APIError `json:"error,omitempty"`
}

type APIError struct {
	Type      ErrorType         `json:"type"`
	Message   string            `json:"message"`
	Details   map[string]string `json:"details,omitempty"`
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
		Type:      errType,
		Message:   message,
		Details:   details,
	}

	json.NewEncoder(w).Encode(APIResponse{
		Data:  nil,
		Error: &err,
	})
}

func WriteServerError(w http.ResponseWriter) {
	WriteError(w, http.StatusInternalServerError, ErrInternal, "Internal Server Error. Please try again later.", nil)	
}
