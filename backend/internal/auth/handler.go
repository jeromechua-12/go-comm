package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jeromechua-12/go-comm/api"
)

type Handler struct {
	svc *service
}

func NewHandler(s *service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) UserSignup(w http.ResponseWriter, r *http.Request) {
	// decode json form data
	var form UserSignupForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, api.ErrBadRequest, "Bad Request", nil)
		return
	}

	// validate fields
	fieldErrors := validateSignup(form)

	if len(fieldErrors) > 0 {
		api.WriteError(w, http.StatusUnprocessableEntity, api.ErrValidation, "Validation Failed", fieldErrors)
		return
	}

	// create new user
	err = h.svc.Signup(r.Context(), form)
	if err != nil {
		// duplicate email
		if errors.Is(err, ErrDuplicateEmail) {
			fieldErrors["email"] = "Email address is already in use"
			api.WriteError(w, http.StatusUnprocessableEntity, api.ErrValidation, "Validation Failed", fieldErrors)
			return
		} else {
			api.WriteError(
				w,
				http.StatusInternalServerError,
				api.ErrInternal,
				"Internal Server Error. Please try again later.",
				nil,
			)
			return
		}
	}

	// no error, send successful response
	api.WriteSuccess(w, http.StatusCreated, "User successfully created")
}
