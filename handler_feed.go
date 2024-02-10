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

func handlerCreateFeed(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {

	type body struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := body{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
	}

	user := r.Context().Value("user").(database.User)

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't create feed:%v", err))
		return
	}
	utils.RespondWithJson(w, 201, databaseFeed(feed))
}

func handlerGetFeedByUser(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(database.User)
	feeds, err := cfg.DB.GetFeedByUser(r.Context(), user.ID)

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't get feed: %v", err))
		return
	}

	utils.RespondWithJson(w, 200, databaseFeeds(feeds))
}

func handlerGetFeeds(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't get feed: %v", err))
		return
	}

	utils.RespondWithJson(w, 200, databaseFeeds(feeds))
}
