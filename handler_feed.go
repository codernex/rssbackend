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

func (cfg apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

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

func (cfg apiConfig) handlerGetFeedByUser(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := cfg.DB.GetFeedByUser(r.Context(), user.ID)

	feed := make([]Feed, len(feeds))

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't get feed: %v", err))
		return
	}

	for i, data := range feeds {
		feed[i] = Feed(data)
	}

	utils.RespondWithJson(w, 200, feed)
}

func (cfg apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	feed := make([]Feed, len(feeds))

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't get feed: %v", err))
		return
	}

	for i, data := range feeds {
		feed[i] = Feed(data)
	}

	utils.RespondWithJson(w, 200, feed)
}
