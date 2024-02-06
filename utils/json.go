package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errResponse{
		Error: msg,
	})
}
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(200)
	_, e := w.Write(data)

	if e != nil {
		log.Printf("Failed to write response %v", e)
	}

}
