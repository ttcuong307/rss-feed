package main

import (
	"encoding/json"
	"net/http"
	"rss-feed/internal/database"
	model "rss-feed/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		model.User
	}
	type response struct {
		model.User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't decode parameters")
		return
	}

	user := model.CreateUserParams{
		ID:        uuid.NewString(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		ApiKey:    uuid.NewString(),
	}

	err = cfg.DB.CreateUser(r.Context(), user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coundn't create user")
		return
	}

	respondWithJSON(w, http.StatusOK, convUserToUser(user))
}

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
