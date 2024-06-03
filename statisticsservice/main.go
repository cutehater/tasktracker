package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"statisticsservice/controllers"
	"statisticsservice/db"
	"statisticsservice/message_broker"
	"statisticsservice/protos"
)

func main() {
	db.ConnectToDb()
	message_broker.RunConsumer()

	server := grpc.NewServer()
	protos.RegisterStatisticsServiceServer(server, &controllers.StatisticsServiceServer{})

	listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
