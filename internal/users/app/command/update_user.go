package command

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type UpdateUserHandler struct {
	userRepo user.Repository
}

func NewUpdateUserHandler(repo user.Repository) UpdateUserHandler {
	if repo == nil {
		panic("user repo is not provided")
	}

	return UpdateUserHandler{userRepo: repo}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, in *user.UpdateUserInput) error {
	if err := h.userRepo.UpdateUser(ctx, in); err != nil {
		return err
	}

	return nil
}
