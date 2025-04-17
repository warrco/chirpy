package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	response := []Chirp{}

	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to retrieve chirps", err)
		return
	}

	for _, chirp := range chirps {
		response = append(response, Chirp(chirp))
	}

	respondWithJson(w, http.StatusOK, response)
}
