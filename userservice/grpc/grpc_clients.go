package grpc

import (
	"log"
	"os"

	"google.golang.org/grpc"

	"userservice/protos"
)

var GRPCTaskServiceClient protos.TaskServiceClient
var GRPCStatisticsServiceClient protos.StatisticsServiceClient

func CreateGRPCClients() {
	conn, err := grpc.Dial(os.Getenv("TASK_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	GRPCTaskServiceClient = protos.NewTaskServiceClient(conn)

	conn, err = grpc.Dial("STATISTICS_SERVICE_ADDR", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	GRPCTaskServiceClient = protos.NewTaskServiceClient(conn)
}
