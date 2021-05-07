package query

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
)

type UsersHandler struct {
	userRepo user.Repository
}

func NewUsersHandler(repo user.Repository) UsersHandler {
	return UsersHandler{userRepo: repo}
}

func (h UsersHandler) Handle(ctx context.Context) []*user.User {
	return h.userRepo.GetUsers(ctx)
}
