package command

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
)

type UserService interface {
	GetUser(ctx context.Context, userId string) (*article.User, error)
	UserExists(ctx context.Context, userId string) error
}
