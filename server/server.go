package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-chi/chi/v5"
)

type ServerInterface interface {
	Start()
	RegisterRoutes(router *chi.Mux)
	RegisterMiddlewares()
	RegisterApiRoutes()
	HealthCheck(http.ResponseWriter, *http.Request)
}

type Server struct {
	Router *chi.Mux
	Logger *logger.Logger
	Config config.EnvConfig
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

func NewServer(router *chi.Mux, logger *logger.Logger, config config.EnvConfig) ServerInterface {
	return &Server{
		Router: router,
		Logger: logger,
		Config: config,
	}
}

func (s *Server) RegisterMiddlewares() {
	if s.Config.IsDevelopment {
		s.Router.Use(ColorLoggingMiddleware)
	}
}
