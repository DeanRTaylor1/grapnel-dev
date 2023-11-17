package services

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><body><h1>Hello, World!</h1></body></html>"))
}
