package main

import (
	"encoding/json"
	"fmt"
	"github.com/codernex/rssbackend/internal/database"
	"github.com/codernex/rssbackend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func handlerCreateFeedFollows(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(database.User)

	type body struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := body{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't follows feed:%v", err))
		return
	}
	utils.RespondWithJson(w, 201, databaseFeedFollow(feedFollow))
}

func handlerGetFeedFollows(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(database.User)

	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	fmt.Println(feedFollows)
	if err != nil {
		utils.RespondWithErr(w, 404, fmt.Sprintf("Coudn't get feed_follows:%v", err))
		return
	}

	utils.RespondWithJson(w, 200, databaseFeedFollows(feedFollows))
}

func handlerDeleteFeedFollows(cfg *utils.ApiConfig, w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(database.User)

	feedFollowId := chi.URLParam(r, "feedFollowId")

	feedFollowIdUuid, err := uuid.Parse(feedFollowId)
	if err != nil {
		utils.RespondWithErr(w, 404, fmt.Sprintf("Coudn't parse feed_follow_id:%v", err))
		return
	}

	err = cfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowIdUuid,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("Coudn't delete feed_follows :%v", err))
		return
	}
	utils.RespondWithJson(w, 200, struct {
		Message string `json:"message"`
	}{
		Message: "Feed Follows Deleted",
	})
}
