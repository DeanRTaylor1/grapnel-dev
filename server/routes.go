package server

import (
	"net/http"
	"path/filepath"

	"github.com/DeanRTaylor1/deans-site/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeHome(w, r, s.Logger)
	})

	router.Get("/styles/output.css", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeCss(w, r, s.Logger)
	})

	router.Get("/faq", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeFaq(w, r, s.Logger)
	})

	router.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeAbout(w, r, s.Logger)
	})

	router.Get("/fonts/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeFonts(w, r, s.Logger)
	})

	// Serve static files
	staticDir := getStaticDir()
	staticFileServer := http.FileServer(http.Dir(staticDir))
	staticRoute := "/static/"

	router.Handle(staticRoute+"*", http.StripPrefix(staticRoute, staticFileServer))
}

func getStaticDir() string {
	dir, err := filepath.Abs(filepath.Join(filepath.Dir("."), "static"))
	if err != nil {
		// Handle error
		panic(err)
	}
	return dir
}
