package main

import (
	"encoding/json"
	"net/http"
	"rss-feed/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't decode parameters")
		return
	}

	feed := database.CreateFeedParams{
		ID:        uuid.NewString(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = cfg.DB.CreateFeed(r.Context(), feed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't create feed")
		return
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.NewString(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = cfg.DB.CreateFeedFollow(r.Context(), feedFollow)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		feed       Feed
		feedFollow FeedFollow
	}{
		feed:       convFeedToFeed(feed),
		feedFollow: convFeedFollowToFeedFollow(feedFollow),
	})
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Feeds []database.Feed `json:"feeds"`
	}

	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, response{Feeds: feeds})
}
