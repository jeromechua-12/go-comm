package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jeromechua-12/go-comm/internal"
)

type handler struct {
	svc *service
}

func newHandler(s *service) *handler {
	return &handler{svc: s}
}

func (h *handler) userSignup(w http.ResponseWriter, r *http.Request) {
	// decode json form data
	var form UserSignupForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		internal.WriteError(w, http.StatusBadRequest, internal.ErrBadRequest, "Bad Request", nil)
		return
	}

	// validate fields
	fieldErrors := validateSignup(form)

	if len(fieldErrors) > 0 {
		internal.WriteError(w, http.StatusUnprocessableEntity, internal.ErrValidation, "Validation Failed", fieldErrors)
		return
	}

	// create new user
	err = h.svc.signup(r.Context(), form)
	if err != nil {
		// duplicate email
		if errors.Is(err, ErrDuplicateEmail) {
			fieldErrors["email"] = "Email address is already in use"
			internal.WriteError(w, http.StatusUnprocessableEntity, internal.ErrValidation, "Validation Failed", fieldErrors)
			return
		}
		internal.WriteServerError(w)
		return
	}

	// no error, send successful response
	internal.WriteSuccess(w, http.StatusCreated, "User successfully created")
}

func (h *handler) userLogin(w http.ResponseWriter, r *http.Request) {
	// decode json form data
	var form UserLoginForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		internal.WriteError(w, http.StatusBadRequest, internal.ErrBadRequest, "Bad Request", nil)
		return
	}

	authRes, err := h.svc.authenticate(r.Context(), form)
	if err != nil {
		if errors.Is(err, ErrBadCredentials) {
			internal.WriteError(w, http.StatusUnauthorized, internal.ErrBadCredentials, "Invalid Email or Password", nil)
			return
		}
		internal.WriteServerError(w)
		return
	}

	accessCookie := http.Cookie{
		Name: "access_token",
		Value: authRes.AccessToken,
		Path: "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge: int(1 * time.Hour / time.Second),
	}

	http.SetCookie(w, &accessCookie)
	internal.WriteSuccess(w, http.StatusCreated, authRes.UserInfo)
}
