package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerInterface interface {
	Start()
	RegisterRoutes(router *chi.Mux)
	RegisterMiddlewares()
	RegisterApiRoutes()
	HealthCheck(http.ResponseWriter, *http.Request)
}

type Server struct {
	Router      *chi.Mux
	Logger      *logger.Logger
	Config      config.EnvConfig
	MongoClient *mongo.Client
}

func (s *Server) Start() {
	s.logStartupMessage()

	// Set up the HTTP server
	srv := &http.Server{
		Addr:    ":8080", // Replace with your server's address and port
		Handler: s.Router,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	s.Logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	// Disconnect MongoDB
	if err := s.MongoClient.Disconnect(ctx); err != nil {
		s.Logger.Error(fmt.Sprintf("Failed to disconnect MongoDB: %v", err))
	}

	s.Logger.Info("Server has shut down")
}

func NewServer(router *chi.Mux, logger *logger.Logger, config config.EnvConfig) ServerInterface {
	mongoClient, err := connectMongoDB()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
		os.Exit(1)
	}

	return &Server{
		Router:      router,
		Logger:      logger,
		Config:      config,
		MongoClient: mongoClient,
	}
}

func connectMongoDB() (*mongo.Client, error) {
	uri := config.Env.Mongo_Uri // Replace with your MongoDB URI
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *Server) RegisterMiddlewares() {
	if s.Config.IsDevelopment {
		s.Router.Use(ColorLoggingMiddleware)
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
