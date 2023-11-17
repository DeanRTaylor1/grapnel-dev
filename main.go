package main

import (
	"fmt"

	"github.com/DeanRTaylor1/deans-site/server"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Hello world")

	r := chi.NewRouter()

	s := server.NewServer(r)
	s.RegisterRoutes(r)

	s.Start()
}
