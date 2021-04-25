package query

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
)

type ArticlesHandler struct {
	articleRepo article.Repository
}

func NewArticlesHandler(repo article.Repository) ArticlesHandler {
	if repo == nil {
		panic("nil article repository")
	}

	return ArticlesHandler{articleRepo: repo}
}

func (handler ArticlesHandler) Handle(ctx context.Context) ([]*article.Article, error) {
	result, err := handler.articleRepo.GetArticles(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
