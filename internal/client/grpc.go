package client

import (
	"errors"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	articlePb "0AlexZhong0/goblog/internal/generated/api/protobuf/article_service"
	userPb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"
)

func NewUserClient() (client userPb.UserServiceClient, close func() error, err error) {
	grpcAddr := viper.GetString("USERS_GRPC_ADDR")

	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env USERS_GRPC_ADDR")
	}

	opts := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return userPb.NewUserServiceClient(conn), conn.Close, nil
}

func NewArticleClient() (client articlePb.ArticleServiceClient, close func() error, err error) {
	grpcAddr := viper.GetString("ARTICLES_GRPC_ADDR")

	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env ARTICLES_GRPC_ADDR")
	}

	opts := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return articlePb.NewArticleServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(_addr string) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
	}
}
