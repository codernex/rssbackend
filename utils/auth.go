package utils

import (
	"context"
	"fmt"
	"github.com/codernex/rssbackend/internal/auth"
	"net/http"
)

// IsAuthenticated is a middleware function that checks for the presence and validity of an API key in the Authorization header of an HTTP request.
// If the API key is not found or is malformed, it responds with an error message.
// If the API key is valid, it calls the next HTTP handler in the chain.
// Example Usage:
//
//	v1Router.Route("/protected", func(r chi.Router) {
//		r.Use(utils.IsAuthenticated)
//		r.Post("/", apiCfg.handlerCreateUser)
//		r.Get("/", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
//	})
func (cfg ApiConfig) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			apiKey, err := auth.GetAPIKey(r.Header)
			if err != nil {
				RespondWithErr(w, 401, fmt.Sprintf("Auth error: %v", err))
				return
			}
			user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)

			if err != nil {
				RespondWithErr(w, 401, fmt.Sprintf("User Not Found: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
