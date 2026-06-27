package gateway

import (
	"net/http"

	"github.com/jeromechua-12/go-comm/internal/middleware"

	"github.com/justinas/alice"
)

type Registrar interface {
	RegisterRoutes(*http.ServeMux)
}

func NewRouter(registrars ...Registrar) http.Handler {
	mux := http.NewServeMux()

	// add routes
	for _, r := range registrars {
		r.RegisterRoutes(mux)
	}

	// middlewares
	chain := alice.New(middleware.CORSMiddleware)

	return chain.Then(mux)
}
