package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/warrco/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds *int   `json:"expires_in_seconds"`
	}

	type response struct {
		User
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unexepected format in the request", err)
	}

	user, err := cfg.db.GetUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Username or password incorrect", err)
		return
	}

	err = auth.CheckPasswordHash(user.HashedPassword, params.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Username or password incorrect", err)
		return
	}

	expirationSeconds := 3600
	if params.ExpiresInSeconds != nil && *params.ExpiresInSeconds > 0 {
		expirationSeconds = min(*params.ExpiresInSeconds, 3600)
	}

	token, err := auth.MakeJWT(user.ID, cfg.secret, time.Duration(expirationSeconds)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create JWT", err)
		return
	}

	respondWithJson(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token: token,
	})
}
