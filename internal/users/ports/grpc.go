package ports

import (
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"
	"0AlexZhong0/goblog/internal/users/app"
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserGrpcServer struct {
	app app.Application
	pb.UnimplementedUserServiceServer
}

func newPbUserFromUser(u *user.User) *pb.User {
	return &pb.User{
		Id:        u.Id(),
		Avatar:    u.Avatar(),
		FirstName: u.FirstName(),
		LastName:  u.LastName(),
	}
}

func NewUserGrpcServer(app app.Application) *UserGrpcServer {
	return &UserGrpcServer{app: app}
}

func (g *UserGrpcServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	result, err := g.app.GetUser.Handle(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return newPbUserFromUser(result), nil
}

func (g *UserGrpcServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*wrapperspb.BoolValue, error) {
	err := g.app.Commands.CreateUser.Handle(ctx, &user.NewUserInput{
		Id:        in.Id,
		Avatar:    in.Avatar,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	})

	if err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (g *UserGrpcServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*wrapperspb.BoolValue, error) {
	err := g.app.UpdateUser.Handle(ctx, &user.UpdateUserInput{
		Id:        in.Id,
		Avatar:    in.Avatar,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	})

	if err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (g *UserGrpcServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*wrapperspb.BoolValue, error) {
	err := g.app.Commands.DeleteUser.Handle(ctx, in.Id)

	if err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (g *UserGrpcServer) UserExists(ctx context.Context, in *pb.UserExistsRequest) (*wrapperspb.BoolValue, error) {
	err := g.app.Queries.UserExists.Handle(ctx, in.Id)
	if err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (g *UserGrpcServer) GetUsers(_ *emptypb.Empty, stream pb.UserService_GetUsersServer) error {
	users := g.app.Queries.GetUsers.Handle(context.TODO())

	for _, user := range users {
		if err := stream.Send(newPbUserFromUser(user)); err != nil {
			return err
		}
	}

	return nil
}
