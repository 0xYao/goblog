package app

import (
	"0AlexZhong0/goblog/internal/users/adapters"
	"0AlexZhong0/goblog/internal/users/app/command"
	"0AlexZhong0/goblog/internal/users/app/query"
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type Application struct {
	Queries
	Commands
}

type Commands struct {
	DeleteUser command.DeleteUserHandler
	UpdateUser command.UpdateUserHandler
	CreateUser command.CreateUserHandler
}

type Queries struct {
	GetUser query.UserHandler
}

func NewApplication(ctx context.Context) Application {
	userFactory, err := user.NewFactory()
	if err != nil {
		panic(err)
	}

	userRepo := adapters.NewMemoryUserRepository(userFactory)

	return Application{
		Queries: Queries{
			GetUser: query.NewUserHandler(userRepo),
		},
		Commands: Commands{
			CreateUser: command.NewCreateUserHandler(userRepo),
			DeleteUser: command.NewDeleteUserHandler(userRepo),
			UpdateUser: command.NewUpdateUserHandler(userRepo),
		},
	}
}
