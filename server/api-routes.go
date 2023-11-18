package server

import (
	"net/http"

	"github.com/DeanRTaylor1/deans-site/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterApiRoutes() {
	apiRouter := chi.NewRouter()
	apiRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			s.HealthCheck(w, r)
		})

		r.Post("/subscribe", func(w http.ResponseWriter, r *http.Request) {
			handlers.Subscribe(w, r, *s.Logger)
		})
	})

	s.Router.Mount("/", apiRouter)
}
