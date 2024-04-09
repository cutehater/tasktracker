package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"taskservice/controllers"
	"taskservice/db"
	"taskservice/protos"
)

func main() {
	db.ConnectToDb()

	server := grpc.NewServer()

	protos.RegisterTaskServiceServer(server, &controllers.TaskServiceServer{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
