package command

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type DeleteUserHandler struct {
	userRepo user.Repository
}

func NewDeleteUserHandler(repo user.Repository) DeleteUserHandler {
	if repo == nil {
		panic("user repo is not provided")
	}

	return DeleteUserHandler{userRepo: repo}
}

func (h *DeleteUserHandler) Handle(ctx context.Context, userId string) error {
	if err := h.userRepo.DeleteUser(ctx, userId); err != nil {
		return err
	}

	return nil
}
