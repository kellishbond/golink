package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// var store = map[string]string{}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := make([]rune, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	code := generateCode()
	store[code] = body.URL

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"code":      code,
		"short_url": "https://golink-production-8f50.up.railway.app/health/" + code,
	})
}
