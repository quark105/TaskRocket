package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/config"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/task"
	log "github.com/sebastianpacuk/taskrocket/utils/logger"
)

const readHeaderTimeout = 3 * time.Second

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.ServiceLogger.Fatalln("Loading config failed in api gateway service, err:", err)
	}
	log.ConfigureLogger(c.LoggingConfig)

	r := chi.NewRouter()
	task.RegisterRoutes(r, &c)

	server := &http.Server{
		Addr:              ":" + c.Port,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.ServiceLogger.WithField("port", c.Port).Info("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
