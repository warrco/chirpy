package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve chirps", err)
		return
	}

	sortString := r.URL.Query().Get("sort")
	if sortString != "asc" && sortString != "desc" {
		respondWithError(w, http.StatusBadRequest, "Invalid sort method", err)
		return
	}

	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}

	response := []Chirp{}
	for _, chirp := range chirps {
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}

		response = append(response, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	if sortString == "desc" {
		sort.Slice(response, func(i, j int) bool {
			return response[i].CreatedAt.After(response[j].CreatedAt)
		})
	}

	respondWithJson(w, http.StatusOK, response)
}

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "failed to get chirp", err)
		return
	}

	respondWithJson(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
