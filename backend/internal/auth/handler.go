package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jeromechua-12/go-comm/api"
)

type Handler struct {
	svc *service
}

func NewHandler(s *service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/user/signup", h.UserSignup)
	mux.HandleFunc("POST /api/user/login", h.UserLogin)
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
		}
		api.WriteServerError(w)
		return
	}

	// no error, send successful response
	api.WriteSuccess(w, http.StatusCreated, "User successfully created")
}

func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	// decode json form data
	var form UserLoginForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, api.ErrBadRequest, "Bad Request", nil)
		return
	}

	authRes, err := h.svc.Authenticate(r.Context(), form)
	if err != nil {
		if errors.Is(err, ErrBadCredentials) {
			api.WriteError(w, http.StatusUnauthorized, api.ErrBadCredentials, "Invalid Email or Password", nil)
		}
		api.WriteServerError(w)
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
	api.WriteSuccess(w, http.StatusCreated, authRes.UserInfo)
}
