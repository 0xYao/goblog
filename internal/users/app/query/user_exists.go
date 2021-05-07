package query

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type UserExistsHandler struct {
	userRepo user.Repository
}

func NewUserExistsHandler(repo user.Repository) UserExistsHandler {
	if repo == nil {
		panic("user repo is not provided")
	}

	return UserExistsHandler{userRepo: repo}
}

func (h *UserExistsHandler) Handle(ctx context.Context, userId string) error {
	return h.userRepo.UserExists(ctx, userId)
}
