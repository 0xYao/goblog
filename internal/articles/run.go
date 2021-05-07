package articles

import (
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

func Run() {
	ctx := context.Background()
	application, cleanUpConnections := app.NewApplication(ctx)
	port := viper.GetString("ARTICLES_GRPC_PORT")
	serverType := strings.ToLower(viper.GetString("SERVER_TYPE"))

	defer cleanUpConnections()

	switch serverType {
	case "grpc":
		server.RunGrpcServer(port, func(server *grpc.Server) {
			grpcServer := ports.NewGrpcServer(application)
			pb.RegisterArticleServiceServer(server, grpcServer)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
