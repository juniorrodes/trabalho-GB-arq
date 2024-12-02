package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	logger = log.New(os.Stdout, "[Sample server]", log.Ldate|log.Ltime|log.Lshortfile)
)

type ResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

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
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := ResponseWriter{ResponseWriter: w}
		next.ServeHTTP(&rw, r)

		logger.Printf("Completed request in %v and with status %d", time.Since(start), rw.statusCode)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("POST /", logRequest(http.HandlerFunc(echo)))
	mux.HandleFunc("GET /health", health)

	logger.Println("Listening at :8080")
	logger.Fatal(http.ListenAndServe(":8080", mux).Error())
}
