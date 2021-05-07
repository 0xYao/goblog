package adapters

import (
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
	"errors"
	"sync"
)

type MemoryArticleRepository struct {
	articles map[string]*article.Article
	lock     *sync.RWMutex

	articleFactory article.Factory
}

func (m *MemoryArticleRepository) getValuesFromArticleMap() []*article.Article {
	articleArray := make([]*article.Article, len(m.articles))

	for _, a := range m.articles {
		articleArray = append(articleArray, a)
	}

	return articleArray
}

func NewMemoryArticleRepository(articleFactory article.Factory) *MemoryArticleRepository {
	return &MemoryArticleRepository{
		lock:           &sync.RWMutex{},
		articleFactory: articleFactory,
		articles:       map[string]*article.Article{},
	}
}

func (m *MemoryArticleRepository) GetArticle(ctx context.Context, articleId string) (*article.Article, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result, exists := m.articles[articleId]
	if !exists {
		return nil, errors.New("article does not exist")
	}

	return result, nil
}

func (m *MemoryArticleRepository) GetArticles(ctx context.Context) ([]*article.Article, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := []*article.Article{}

	for _, article := range m.articles {
		if !article.IsDraft() {
			result = append(result, article)
		}
	}

	return result, nil
}

func (m *MemoryArticleRepository) AddArticle(ctx context.Context, in *article.NewArticleInput) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	articleArray := m.getValuesFromArticleMap()

	// validation
	if err := article.IsArticleTitleUnique(in.Title, articleArray); err != nil {
		return err
	}

	newArticle, err := m.articleFactory.NewArticle(in)
	if err != nil {
		return err
	}

	// add the article
	m.articles[newArticle.Id()] = newArticle
	return nil
}
