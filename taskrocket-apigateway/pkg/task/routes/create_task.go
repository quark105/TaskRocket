package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sebastianpacuk/taskrocket/taskrocket-apigateway/pkg/task/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateTaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
	DueDate     string `json:"due_date"`
}

func CreateTask(w http.ResponseWriter, r *http.Request, c pb.TaskServiceClient) {
	body := CreateTaskRequestBody{}

	rbody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(rbody, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := time.Parse(time.DateOnly, body.DueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.CreateTask(context.Background(), &pb.CreateTaskRequest{
		Title:       body.Title,
		Description: body.Description,
		Status:      pb.Status(body.Status),
		DueDate:     timestamppb.New(t),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
