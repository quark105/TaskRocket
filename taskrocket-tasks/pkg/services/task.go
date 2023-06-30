package services

import (
	"context"
	"net/http"

	"github.com/sebastianpacuk/taskrocket/taskrocket-tasks/pkg/db"
	"github.com/sebastianpacuk/taskrocket/taskrocket-tasks/pkg/models"
	pb "github.com/sebastianpacuk/taskrocket/taskrocket-tasks/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedTaskServiceServer
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	var task models.Task

	task.Title = req.Title
	task.Description = req.Description
	task.Status = int8(req.Status)
	task.DueDate = req.DueDate.AsTime()

	if result := s.H.DB.Create(&task); result.Error != nil {
		return &pb.CreateTaskResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateTaskResponse{
		Status: http.StatusCreated,
		Id:     task.Id,
	}, nil
}
