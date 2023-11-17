package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/DeanRTaylor1/deans-site/server"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Can not get current directory")
	}

	l, err := logger.NewLogger(fmt.Sprintf("%s%s", cwd, "/logs"))
	if err != nil {
		log.Fatal("Unable to initialise logger.")
	}

	s := server.NewServer(r, l)
	s.RegisterMiddlewares()
	s.RegisterRoutes(r)

	s.Start()
}
