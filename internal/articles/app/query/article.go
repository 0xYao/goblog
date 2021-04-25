package query

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
)

type ArticleHandler struct {
	articleRepo article.Repository
}

func NewArticleHandler(repo article.Repository) ArticleHandler {
	if repo == nil {
		panic("nil article repository")
	}

	return ArticleHandler{articleRepo: repo}
}

func (handler ArticleHandler) Handle(ctx context.Context, articleId string) (*article.Article, error) {
	result, err := handler.articleRepo.GetArticle(ctx, articleId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
