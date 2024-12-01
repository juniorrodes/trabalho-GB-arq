package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	logger = log.New(os.Stdout, "[Sample server]", log.Ldate|log.Ltime|log.Lshortfile)
)

func echo(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Expected 'content-type' header, received nothing\n"))
		return
	}

	w.Header().Set("Content-Type", contentType)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to read body\n"))
	}
	w.Write(body)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", echo)
	mux.HandleFunc("GET /health", health)

	logger.Println("Listening at :8080")
	logger.Fatal(http.ListenAndServe(":8080", mux).Error())
}
