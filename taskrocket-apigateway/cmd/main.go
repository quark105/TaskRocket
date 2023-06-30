package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/config"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/task"
)

const readHeaderTimeout = 3 * time.Second

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := chi.NewRouter()
	task.RegisterRoutes(r, &c)

	server := &http.Server{
		Addr:              c.Port,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.Println("Starting server on port", c.Port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
