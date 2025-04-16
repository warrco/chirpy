package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedText := profanityFilter(params.Body)

	respondWithJson(w, http.StatusOK, returnVals{
		CleanedBody: cleanedText,
	})
}

func profanityFilter(msg string) string {
	var badwords = []string{"kerfuffle", "sharbert", "fornax"}

	msgSlice := strings.Fields(msg)

	for i, word := range msgSlice {

		for _, badword := range badwords {
			if strings.ToLower(word) == badword {
				censored := "****"
				msgSlice[i] = censored
				break
			}
		}
	}
	msg = strings.Join(msgSlice, " ")

	return msg
}
