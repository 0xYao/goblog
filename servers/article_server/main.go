package main

import (
	pb "0AlexZhong0/goblog/proto/article_service"
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	pb.UnimplementedArticleServiceServer
	articleMap map[string]pb.Article
}

const (
	_excerptLength = 100
)

func createArticleFromWriteRequest(in *pb.WriteArticleRequest, isDraft bool) pb.Article {
	return pb.Article{
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

func getArticleSummary(article pb.Article) *pb.ArticleSummary {
	return &pb.ArticleSummary{
		Title:      article.Title,
		CoverImage: article.CoverImage,
		Excerpt:    createExcerptFromBody(article.Body),
	}
}

func (s *server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.Article, error) {
	article, exists := s.articleMap[in.ArticleId]

	if !exists {
		return nil, status.Errorf(codes.NotFound, "Article is not found")
	}

	return &article, nil
}

func (s *server) WriteArticle(ctx context.Context, in *pb.WriteArticleRequest) (*wrapperspb.BoolValue, error) {
	for _, article := range s.articleMap {
		if article.Title == in.Title {
			return &wrapperspb.BoolValue{Value: false}, status.Errorf(codes.AlreadyExists, fmt.Sprintf("Article with title %v already exists", in.Title))
		}
	}

	newArticle := createArticleFromWriteRequest(in, false)
	s.articleMap[newArticle.Id] = newArticle

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (s *server) SaveArticleAsDraft(ctx context.Context, in *pb.WriteArticleRequest) (*wrapperspb.BoolValue, error) {
	articleDraft := createArticleFromWriteRequest(in, true)
	s.articleMap[articleDraft.Id] = articleDraft
	return &wrapperspb.BoolValue{Value: true}, nil
}

func (s *server) GetArticles(_ *emptypb.Empty, stream pb.ArticleService_GetArticlesServer) error {
	for _, article := range s.articleMap {
		if !article.IsDraft {
			if err := stream.Send(getArticleSummary(article)); err != nil {
				return err
			}
		}
	}

	return nil
}

func newServer() *server {
	// init and load the data here
	s := &server{articleMap: make(map[string]pb.Article)}
	return s
}

func main() {
	addr := fmt.Sprintf("localhost:%d", 8080)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterArticleServiceServer(grpcServer, newServer())

	log.Printf("Server is running at: %v", addr)
	grpcServer.Serve(lis)
}
