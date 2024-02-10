package main

import (
	"github.com/codernex/rssbackend/utils"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct {
	}{})

}
