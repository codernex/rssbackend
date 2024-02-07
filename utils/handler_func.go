package utils

import "net/http"

type HandlerFunc func(cfg *ApiConfig, w http.ResponseWriter, r *http.Request)

func (cfg ApiConfig) RequestHandler(handlerFunc HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		handlerFunc(&cfg, writer, request)
	}
}
