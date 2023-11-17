package server

import (
	"fmt"
	"net/http"
	"os"

	services "github.com/DeanRTaylor1/deans-site/internal"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-chi/chi"
)

type ServerInterface interface {
	Start()
	RegisterRoutes(router *chi.Mux)
	RegisterMiddlewares()
}

type Server struct {
	Router *chi.Mux
	Logger *logger.Logger
}

func (s *Server) Start() {
	// Log startup message
	s.logStartupMessage()

	// Start the server
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		os.Exit(1)
	}
}

func (s *Server) logStartupMessage() {
	art := `
    ██████╗ ███████╗ █████╗ ██████╗ ██╗   ██╗
    ██╔══██╗██╔════╝██╔══██╗██╔══██╗╚██╗ ██╔╝
    ██████╔╝█████╗  ███████║██║  ██║ ╚████╔╝
    ██╔══██╗██╔══╝  ██╔══██║██║  ██║  ╚██╔╝
    ██║  ██║███████╗██║  ██║██████╔╝   ██║
    ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝    ╚═╝
  `

	s.Logger.Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	s.Logger.Info(art)
	s.Logger.Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	s.Logger.Info("🚀 Application is running on: http://localhost:8080/")
	s.Logger.Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	s.Logger.Info("")
}

func NewServer(router *chi.Mux, logger *logger.Logger) ServerInterface {
	return &Server{
		Router: router,
		Logger: logger,
	}
}

func (s *Server) RegisterMiddlewares() {
	s.Router.Use(ColorLoggingMiddleware)
}

func (s *Server) RegisterRoutes(router *chi.Mux) {
	// s.Router.Route("/api/v1", func(r chi.Router) {
	// 	r.Mount("/auth", authController.Routes())
	// })

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		services.Home(w, r, s.Logger)

	})
}
