package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sebastianpacuk/taskrocket-tasks/pkg/config"
	"github.com/sebastianpacuk/taskrocket-tasks/pkg/db"
	pb "github.com/sebastianpacuk/taskrocket-tasks/pkg/pb"
	services "github.com/sebastianpacuk/taskrocket-tasks/pkg/services"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Task Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTaskServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
