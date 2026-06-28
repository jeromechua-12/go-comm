package middleware

import (
	"log/slog"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3030")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var (
				ip = r.RemoteAddr
				proto = r.Proto
				method = r.Method
				uri = r.URL.RequestURI()
			)

			logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn) 
	}
}
