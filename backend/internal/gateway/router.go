package gateway

import (
	"log/slog"
	"net/http"

	"github.com/jeromechua-12/go-comm/internal/middleware"

	"github.com/justinas/alice"
)

type Registrar interface {
	RegisterRoutes(*http.ServeMux)
}

func NewRouter(logger *slog.Logger, registrars ...Registrar) http.Handler {
	mux := http.NewServeMux()

	// add routes
	for _, r := range registrars {
		r.RegisterRoutes(mux)
	}

	// middlewares
	chain := alice.New(
		middleware.LoggingMiddleware(logger),
		middleware.CORSMiddleware,
	)

	return chain.Then(mux)
}
