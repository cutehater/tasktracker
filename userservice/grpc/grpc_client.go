package grpc

import (
	"log"

	"google.golang.org/grpc"

	"userservice/protos"
)

var GRPCClient protos.TaskServiceClient

func CreateGRPCClient() {
	conn, err := grpc.Dial("taskservice:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	GRPCClient = protos.NewTaskServiceClient(conn)
}
