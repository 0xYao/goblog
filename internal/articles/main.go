package main

import (
	"0AlexZhong0/goblog/config"
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/article_service"

	"context"
	"fmt"
	"strings"

	"0AlexZhong0/goblog/internal/articles/app"
	"0AlexZhong0/goblog/internal/articles/ports"
	"0AlexZhong0/goblog/internal/common/server"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	ctx := context.Background()
	application := app.NewApplication(ctx)
	serverType := strings.ToLower(viper.GetString("SERVER_TYPE"))

	switch serverType {
	case "grpc":
		server.RunGrpcServer(func(server *grpc.Server) {
			grpcServer := ports.NewGrpcServer(application)
			pb.RegisterArticleServiceServer(server, grpcServer)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
