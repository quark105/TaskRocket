package task

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sebastianpacuk/taskrocket-apigateway/pkg/config"
	"github.com/sebastianpacuk/taskrocket-apigateway/pkg/task/pb"
)

type ServiceClient struct {
	Client pb.TaskServiceClient
}

func InitTaskServiceClient(c *config.Config) pb.TaskServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.TaskSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewTaskServiceClient(cc)
}
