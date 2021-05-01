package user

import "context"

type Repository interface {
	DeleteUser(ctx context.Context, userId string) (error)
	CreateUser(ctx context.Context, in *NewUserInput) (error)
	GetUser(ctx context.Context, userId string) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserInput) (error)
}
