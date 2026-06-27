package gateway

import (
	"net/http"

	"github.com/jeromechua-12/go-comm/internal/auth"

	"github.com/justinas/alice"
)

func NewRouter(authHandler *auth.Handler) http.Handler {
	mux := http.NewServeMux()

	// add routes
	mux.HandleFunc("POST /api/user/signup", authHandler.UserSignup)
	mux.HandleFunc("POST /api/user/login", authHandler.UserLogin)

	// middlewares
	chain := alice.New(CORSMiddleware)

	return chain.Then(mux)
}
