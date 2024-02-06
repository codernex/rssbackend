package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			val := r.Header.Get("Authorization")

			if val == "" {
				RespondWithErr(w, 400, fmt.Sprintf("Api Key Not Found"))
				return
			}
			tokens := strings.Split(val, " ")
			if len(tokens) != 2 {
				RespondWithErr(w, 400, fmt.Sprintf("Api Key Mailformed"))
				return
			}

			if tokens[0] != "ApiKey" {
				RespondWithErr(w, 400, fmt.Sprintf("Api Key Mailformed"))
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
