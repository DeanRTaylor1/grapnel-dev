package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) RegisterApiRoutes() {
	apiRouter := chi.NewRouter()
	apiRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			s.HealthCheck(w, r)
		})
	})

	s.Router.Mount("/", apiRouter)
}
