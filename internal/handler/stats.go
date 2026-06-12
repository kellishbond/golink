package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/stats/")

	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	originalURL, exists := store[code]
	if !exists {
		http.Error(w, "short link not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"code":   code,
		"url":    originalURL,
		"clicks": clicks[code],
	})
}
