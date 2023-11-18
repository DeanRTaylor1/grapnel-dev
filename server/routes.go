package server

import (
	"net/http"

	services "github.com/DeanRTaylor1/deans-site/internal"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		services.Home(w, r, s.Logger)

	})
}
