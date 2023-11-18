package server

import (
	"net/http"

	"github.com/DeanRTaylor1/deans-site/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeHome(w, r, s.Logger)

	})
}
