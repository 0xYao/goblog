package command

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
)

type AddArticlesHandler struct {
	articleRepo article.Repository
}

func NewAddArticlesHandler(repo article.Repository) AddArticlesHandler {
	if repo == nil {
		panic("article repository is not provided")
	}

	return AddArticlesHandler{articleRepo: repo}
}

func (handler AddArticlesHandler) Handle(ctx context.Context, articles []*article.NewArticleInput) error {
	for _, articleInput := range articles {
		if err := handler.articleRepo.AddArticle(ctx, articleInput); err != nil {
			return err
		}
	}

	return nil
}
