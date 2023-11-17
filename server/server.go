package server

import (
	"fmt"
	"net/http"

	services "github.com/DeanRTaylor1/deans-site/internal"
	"github.com/go-chi/chi"
)

type ServerInterface interface {
	Start()
	RegisterRoutes(router *chi.Mux)
}

type Server struct {
	Router *chi.Mux
}

func (s *Server) Start() {

	http.ListenAndServe(fmt.Sprintf(":%s", "8080"), s.Router)
}

func NewServer(router *chi.Mux) ServerInterface {
	return &Server{
		Router: router,
	}
}

// func (s *Server) RegisterMiddlewares(authenticator authentication.Authenticator, store db.Store) {
// 	s.Router.Use(middleware.ColorLoggingMiddleware)
// 	s.Router.Use(middleware.AuthMiddleware(authenticator, store))
// }

func (s *Server) RegisterRoutes(router *chi.Mux) {
	// s.Router.Route("/api/v1", func(r chi.Router) {
	// 	r.Mount("/auth", authController.Routes())
	// })

	router.Get("/", services.Home)
}
