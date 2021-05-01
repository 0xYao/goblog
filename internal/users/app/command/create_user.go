package command

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type CreateUserHandler struct {
	userRepo user.Repository
}

func NewCreateUserHandler(repo user.Repository) CreateUserHandler {
	if repo == nil {
		panic("user repo is not provided")
	}

	return CreateUserHandler{userRepo: repo}
}

func (h *CreateUserHandler) Handle(ctx context.Context, in *user.NewUserInput) error {
	if err := h.userRepo.CreateUser(ctx, in); err != nil {
		return err
	}

	return nil
}
