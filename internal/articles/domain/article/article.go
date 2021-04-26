package article

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// the fields are private to the current package
type Article struct {
	id         string
	body       string
	title      string
	isDraft    bool
	coverImage string
	createdAt  time.Time
}

// getters
func (a *Article) Id() string {
	return a.id
}

func (a *Article) IsDraft() bool {
	return a.isDraft
}

func (a *Article) Title() string {
	return a.title
}

func (a *Article) CoverImage() string {
	return a.coverImage
}

func (a *Article) Body() string {
	return a.body
}

func (a *Article) CreatedAt() time.Time {
	return a.createdAt
}

// factory encapsulates the logic of creating domain objects
type Factory struct{}

func NewFactory() (Factory, error) {
	return Factory{}, nil
}

type NewArticleInput struct {
	Body       string
	Title      string
	IsDraft    bool
	CoverImage string
}

func (f Factory) NewArticle(in *NewArticleInput) (*Article, error) {
	if in.Body == "" {
		return nil, errors.New("article body is empty")
	}

	if in.Title == "" {
		return nil, errors.New("article title is empty")
	}

	if in.CoverImage == "" {
		return nil, errors.New("article cover image is empty")
	}

	return &Article{
		id:         uuid.NewString(),
		body:       in.Body,
		title:      in.Title,
		isDraft:    in.IsDraft,
		coverImage: in.CoverImage,
		createdAt:  time.Now(),
	}, nil
}
