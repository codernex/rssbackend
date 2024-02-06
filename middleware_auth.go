package main

import (
	"fmt"
	"github.com/codernex/rssbackend/internal/auth"
	"github.com/codernex/rssbackend/internal/database"
	"github.com/codernex/rssbackend/utils"
	"net/http"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithErr(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
