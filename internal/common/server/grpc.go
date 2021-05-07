package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGrpcServer(port string, registerServer func(server *grpc.Server)) {
	addr := fmt.Sprintf(":%s", port)
	RunGrpcServerOnAddr(addr, registerServer)
}

func RunGrpcServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer()

	registerServer(grpcServer)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start the listener: %v", err)
	}

	log.Printf("Listening at %v\n", addr)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
