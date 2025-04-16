package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerMetricHits(w http.ResponseWriter, r *http.Request) {
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
					<html>
					<body>
					<h1>Welcome, Chirpy Admin</h1>
					<p>Chirpy has been visited %d times!</p>
					</body>
					</html>`, cfg.fileserverHits.Load())
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlContent))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
