package main

import (
	"net/http"
	"time"

	"github.com/warrco/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Token string `json:"token"`
	}

	refToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "could not retrieve token", err)
		return
	}

	userID, err := cfg.db.GetUserFromRefreshToken(r.Context(), refToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "failed to retrieve user ID", err)
		return
	}

	token, err := auth.MakeJWT(userID.ID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create new access token", err)
		return
	}

	respondWithJson(w, http.StatusOK, response{
		Token: token,
	})
}
