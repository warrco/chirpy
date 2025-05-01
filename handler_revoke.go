package main

import (
	"net/http"

	"github.com/warrco/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {

	refToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not retrieve token", err)
		return
	}

	err = cfg.db.UpdateRefreshToken(r.Context(), refToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
