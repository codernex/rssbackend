package main

import (
	"encoding/json"
	"fmt"
	"github.com/codernex/rssbackend/internal/database"
	"github.com/codernex/rssbackend/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type body struct {
		Name string `Name:"name"`
	}
	params := body{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't create user:", err))
		return
	}
	utils.RespondWithJson(w, 201, databaseUserToUser(user))
}

func (cfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJson(w, 200, databaseUserToUser(user))
}
