package task

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/config"
	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/task/routes"
)

func RegisterRoutes(r *chi.Mux, c *config.Config) {
	svc := &ServiceClient{
		Client: InitTaskServiceClient(c),
	}
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", svc.CreateTask)
	})
}

func (svc *ServiceClient) CreateTask(w http.ResponseWriter, r *http.Request) {
	routes.CreateTask(w, r, svc.Client)
}
