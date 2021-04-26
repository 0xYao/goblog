package ports

import (
	"0AlexZhong0/goblog/internal/articles/app"
	"0AlexZhong0/goblog/internal/articles/domain/article"
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/article_service"
	"context"
	"math"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GrpcServer struct {
	app app.Application
	pb.UnimplementedArticleServiceServer
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app, pb.UnimplementedArticleServiceServer{}}
}

const (
	_excerptLength = 100
)

// helper methods
func newProtoArticle(a *article.Article) *pb.Article {
	return &pb.Article{
		Id:         uuid.NewString(),
		Body:       a.Body(),
		Title:      a.Title(),
		IsDraft:    a.IsDraft(),
		CoverImage: a.CoverImage(),
		CreatedAt:  timestamppb.Now(),
	}
}

func newArticleInputFromProtoWriteReq(in *pb.WriteArticleRequest, isDraft bool) *article.NewArticleInput {
	return &article.NewArticleInput{
		Body:       in.Body,
		Title:      in.Title,
		IsDraft:    isDraft,
		CoverImage: in.CoverImage,
	}
}

func createExcerptFromBody(body string) string {
	length := int32(math.Min(_excerptLength, float64(len(body))))
	return body[:length] + "..."
}

func newArticleSummary(article *article.Article) *pb.ArticleSummary {
	return &pb.ArticleSummary{
		Title:      article.Title(),
		CoverImage: article.CoverImage(),
		Excerpt:    createExcerptFromBody(article.Body()),
	}
}

func (g GrpcServer) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.Article, error) {
	article, err := g.app.GetArticle.Handle(ctx, in.ArticleId)
	if err != nil {
		return nil, err
	}

	return newProtoArticle(article), err
}

func (g GrpcServer) GetArticles(_ *emptypb.Empty, stream pb.ArticleService_GetArticlesServer) error {
	articles, err := g.app.GetArticles.Handle(context.TODO())

	if err != nil {
		return err
	}

	for _, article := range articles {
		if err := stream.Send(newArticleSummary(article)); err != nil {
			return err
		}
	}

	return nil
}

func (g GrpcServer) WriteArticle(ctx context.Context, in *pb.WriteArticleRequest) (*wrapperspb.BoolValue, error) {
	newArticleInput := newArticleInputFromProtoWriteReq(in, false)

	if err := g.app.AddArticles.Handle(ctx, []*article.NewArticleInput{newArticleInput}); err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (g GrpcServer) SaveArticleAsDraft(ctx context.Context, in *pb.WriteArticleRequest) (*wrapperspb.BoolValue, error) {
	newDraftInput := newArticleInputFromProtoWriteReq(in, true)

	if err := g.app.AddArticles.Handle(ctx, []*article.NewArticleInput{newDraftInput}); err != nil {
		return &wrapperspb.BoolValue{Value: false}, err
	}

	return &wrapperspb.BoolValue{Value: true}, nil
}
