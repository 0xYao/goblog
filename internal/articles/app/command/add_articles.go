package command

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
)

type AddArticlesHandler struct {
	articleRepo article.Repository
	userService UserService
}

func NewAddArticlesHandler(repo article.Repository, userService UserService) AddArticlesHandler {
	if repo == nil {
		panic("article repository is not provided")
	}

	if userService == nil {
		panic("user service is not provided")
	}

	return AddArticlesHandler{articleRepo: repo, userService: userService}
}

func (handler AddArticlesHandler) Handle(ctx context.Context, articles []*article.NewArticleInput) error {
	for _, articleInput := range articles {
		if err := handler.userService.UserExists(ctx, articleInput.UserId); err != nil {
			return err
		}

		if err := handler.articleRepo.AddArticle(ctx, articleInput); err != nil {
			return err
		}
	}

	return nil
}
