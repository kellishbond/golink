package main

import (
	"fmt"
	"net/http"

	"github.com/kellishbond/golink/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/", handler.Redirect)

	fmt.Println("GoLink server running on :8080")
	http.ListenAndServe(":8080", mux)
}
