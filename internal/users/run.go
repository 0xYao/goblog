package users

import (
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"

	"0AlexZhong0/goblog/internal/common/server"
	"0AlexZhong0/goblog/internal/users/app"
	"0AlexZhong0/goblog/internal/users/ports"

	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func Run() {
	ctx := context.Background()
	application := app.NewApplication(ctx)
	port := viper.GetString("USERS_GRPC_PORT")

	server.RunGrpcServer(port, func(server *grpc.Server) {
		userGrpcServer := ports.NewUserGrpcServer(application)
		pb.RegisterUserServiceServer(server, userGrpcServer)
	})
}
