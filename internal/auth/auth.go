package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey{insert apikey here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication info found")
	}
	tokens := strings.Split(val, " ")
	if len(tokens) != 2 {
		return "", errors.New("malformed auth header")
	}

	if tokens[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}
	return tokens[1], nil
}
