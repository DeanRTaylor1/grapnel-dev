package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/DeanRTaylor1/deans-site/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()
	config := config.Env

	r := chi.NewRouter()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can not get current directory. Error: %s", err.Error())
	}

	l, err := logger.NewLogger(fmt.Sprintf("%s%s", cwd, "/logs"))
	if err != nil {
		log.Fatalf("Error initialising logger. Error: %s", err.Error())
	}

	s := server.NewServer(r, l, config)
	s.RegisterMiddlewares()
	s.RegisterApiRoutes()
	s.RegisterRoutes(r)

	s.Start()
}
