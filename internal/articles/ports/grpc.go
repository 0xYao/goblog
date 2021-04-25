package main

import (
	pb "0AlexZhong0/goblog/api/protobuf/article_service"
	"0AlexZhong0/goblog/internal/articles/app"
	"0AlexZhong0/goblog/internal/articles/domain/article"
	"context"
	"math"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	// "google.golang.org/protobuf/types/known/emptypb"
	// "google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/protobuf/types/known/wrapperspb"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app}
}

const (
	_excerptLength = 100
)

// helper methods
func newArticleFromWriteRequest(in *pb.WriteArticleRequest, isDraft bool) *pb.Article {
	return &pb.Article{
		Id:         uuid.NewString(),
		Body:       in.Body,
		Title:      in.Title,
		IsDraft:    isDraft,
		CoverImage: in.CoverImage,
		CreatedAt:  timestamppb.Now(),
	}
}

func createExcerptFromBody(body string) string {
	length := int32(math.Min(_excerptLength, float64(len(body))))
	return body[:length] + "..."
}

func getArticleSummary(article article.Article) *pb.ArticleSummary {
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

	return article, err
}
