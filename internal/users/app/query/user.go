package query

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type UserHandler struct {
	userRepo user.Repository
}

func NewUserHandler(repo user.Repository) UserHandler {
	if repo == nil {
		panic("user repo is not provided")
	}

	return UserHandler{userRepo: repo}
}

func (h *UserHandler) Handle(ctx context.Context, userId string) (*user.User, error) {
	user, err := h.userRepo.GetUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
