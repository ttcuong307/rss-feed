package main

import (
	"fmt"
	"net/http"
	"rss-feed/internal/auth"
	"rss-feed/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		fmt.Print(user)
		fmt.Print(err)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
