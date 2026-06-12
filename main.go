package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kellishbond/golink/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/stats/", handler.Stats)
	mux.HandleFunc("/", handler.Redirect)

	fmt.Println("GoLink server running on :" + port)
	http.ListenAndServe(":"+port, mux)
}
