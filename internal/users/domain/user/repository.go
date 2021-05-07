package user

import "context"

type Repository interface {
	GetUsers(ctx context.Context) []*User
	UserExists(ctx context.Context, userId string) error
	DeleteUser(ctx context.Context, userId string) error
	CreateUser(ctx context.Context, in *NewUserInput) error
	GetUser(ctx context.Context, userId string) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserInput) error
}
