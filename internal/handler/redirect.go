package handler

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	originalURL, exists := store[code]
	if !exists {
		http.Error(w, "short link not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
