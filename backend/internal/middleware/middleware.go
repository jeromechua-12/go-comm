package middleware

import (
	"errors"
	"log/slog"
	"os"
	"net/http"

	"github.com/jeromechua-12/go-comm/api"

	"github.com/golang-jwt/jwt/v5"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3030")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

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

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get access token
		cookie, err := r.Cookie("access_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				api.WriteError(w, http.StatusUnauthorized, api.ErrUnauthorized, "Unauthorized", nil)
				return
			}
			api.WriteServerError(w)
		}

		tokenString := cookie.Value

		// parse token
		type CustomClaims struct {
			Role string `json:"role"`
			jwt.RegisteredClaims
		}

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
		})

		// verify signature and expiry
		switch {
		case token.Valid:
			api.WriteSuccess(w, http.StatusOK, nil)	
		case errors.Is(err, jwt.ErrTokenExpired):
			api.WriteError(w, http.StatusUnauthorized, api.ErrAccessExpired, "Access token expired", nil)
		default:
			api.WriteError(w, http.StatusUnauthorized, api.ErrUnauthorized, "Unauthorized", nil)
		}

		next.ServeHTTP(w, r)
	})
}
