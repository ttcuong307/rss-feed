package main

import (
	"encoding/json"
	"net/http"
	"rss-feed/internal/database"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (api *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := api.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID.String(),
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't decode parameters")
		return
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.NewString(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = cfg.DB.CreateFeedFollow(r.Context(), feedFollow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, convFeedFollowToFeedFollow(feedFollow))
}
