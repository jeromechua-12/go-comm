package main 

import (
	"log/slog"
	"net/http"

	"github.com/jeromechua-12/go-comm/internal/middleware"
	"github.com/jeromechua-12/go-comm/module"

	"github.com/justinas/alice"
)

func NewRouter(logger *slog.Logger, modules ...module.Module) http.Handler {
	mux := http.NewServeMux()

	authChain := alice.New(middleware.AuthorizationMiddleware)

	// add routes
	for _, m := range modules {
		for _, r := range m.Routes() {
			routePattern := r.Method + " " + r.Path
			handler := http.Handler(r.HandlerFunc)
			if r.Protected {
				handler = authChain.Then(handler)
			}

			mux.Handle(routePattern, handler)
		}
	}

	// middlewares
	chain := alice.New(
		middleware.LoggingMiddleware(logger),
		middleware.CORSMiddleware,
	)

	return chain.Then(mux)
}
