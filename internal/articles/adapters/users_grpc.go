package adapters

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"
	"context"
)

type UserGrpc struct {
	client pb.UserServiceClient
}

func NewUserGrpc(client pb.UserServiceClient) UserGrpc {
	return UserGrpc{client: client}
}

func unmarshallPbUser(u *pb.User) (*article.User, error) {
	result, err := article.NewUser(article.NewUserInput{
		Id:        u.Id,
		Avatar:    u.Avatar,
		LastName:  u.LastName,
		FirstName: u.FirstName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s UserGrpc) GetUser(ctx context.Context, userId string) (*article.User, error) {
	user, err := s.client.GetUser(ctx, &pb.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}

	unmarshalled, err := unmarshallPbUser(user)
	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func (s UserGrpc) UserExists(ctx context.Context, userId string) error {
	_, err := s.client.UserExists(ctx, &pb.UserExistsRequest{Id: userId})
	if err != nil {
		return err
	}

	return nil
}
