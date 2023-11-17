package server

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/DeanRTaylor1/deans-site/config"
	services "github.com/DeanRTaylor1/deans-site/internal"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-chi/chi"
)

type ServerInterface interface {
	Start()
	RegisterRoutes(router *chi.Mux)
	RegisterMiddlewares()
	RegisterApiRoutes()
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

func (s *Server) RegisterRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		services.Home(w, r, s.Logger)

	})
}

func (s *Server) RegisterApiRoutes() {
	apiRouter := chi.NewRouter()
	apiRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		})
	})

	s.Router.Mount("/", apiRouter)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	cpuUsage := runtime.NumCPU()
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	ramUsage := memStats.HeapAlloc / (1024 * 1024) // in megabytes

	elapsedTime := time.Since(startTime).Milliseconds()

	response := fmt.Sprintf("Server is healthy\nCPU Usage: %d\nRAM Usage: %d MB\nResponse Time: %d ms", cpuUsage, ramUsage, elapsedTime)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(response))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error writing health check response: %s", err))
	}
}
