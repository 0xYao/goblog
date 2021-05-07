package app

import (
	"0AlexZhong0/goblog/internal/articles/adapters"
	"0AlexZhong0/goblog/internal/articles/app/command"
	"0AlexZhong0/goblog/internal/articles/app/query"
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"0AlexZhong0/goblog/internal/client"
	"context"
)

type Application struct {
	Commands
	Queries
}

type Commands struct {
	AddArticles command.AddArticlesHandler
}

type Queries struct {
	GetArticle  query.ArticleHandler
	GetArticles query.ArticlesHandler
}

func NewApplication(ctx context.Context) (Application, func()) {
	articleFactory, err := article.NewFactory()
	if err != nil {
		panic(err)
	}

	// swap out the repo here if we wish to change the database
	articleRepo := adapters.NewMemoryArticleRepository(articleFactory)
	userClient, closeUserClient, err := client.NewUserClient()

	if err != nil {
		panic(err)
	}

	userGrpc := adapters.NewUserGrpc(userClient)

	return Application{
			Commands: Commands{
				AddArticles: command.NewAddArticlesHandler(articleRepo, userGrpc),
			},

			Queries: Queries{
				GetArticle:  query.NewArticleHandler(articleRepo),
				GetArticles: query.NewArticlesHandler(articleRepo),
			},
		}, func() {
			_ = closeUserClient()
		}
}
