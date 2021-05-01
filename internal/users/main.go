package main

import (
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"

	"0AlexZhong0/goblog/config"
	"0AlexZhong0/goblog/internal/common/server"
	"0AlexZhong0/goblog/internal/users/app"
	"0AlexZhong0/goblog/internal/users/ports"

	"context"

	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	ctx := context.Background()
	application := app.NewApplication(ctx)

	server.RunGrpcServer(func(server *grpc.Server) {
		userGrpcServer := ports.NewUserGrpcServer(application)
		pb.RegisterUserServiceServer(server, userGrpcServer)
	})
}
