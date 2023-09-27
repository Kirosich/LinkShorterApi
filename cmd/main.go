package main

import (
	"fmt"
	"go-url-short/internal/transport"
	"net/http"
)

func main() {
	shortener := &transport.URLShortener{
		Urls: make(map[string]string),
	}

	http.HandleFunc("/", shortener.Home)
	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
