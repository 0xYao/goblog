package server

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func RunGrpcServer(registerServer func(server *grpc.Server)) {
	// TODO: use a config management library like viper to handle the config values
	port := viper.GetString("PORT")

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
