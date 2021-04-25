package article

import "context"

type Repository interface {
	AddArticle(ctx context.Context, input *NewArticleInput) error
	GetArticles(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, articleId string) (*Article, error)
}
