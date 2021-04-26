package main

import (
	"flag"
	"io"
	"log"
	"time"

	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/article_service"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	serverAddr = flag.String("server_addr", "localhost:50051", "The server address in the format of host:port")
)

func printArticle(client pb.ArticleServiceClient, articleId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	article, err := client.GetArticle(ctx, &pb.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		log.Fatalf("Failed to get article, %v", err)
	}

	log.Println(article)
}

func printArticles(client pb.ArticleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetArticles(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error getting the articles: %v", err)
	}

	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error receiving an article: %v", err)
		}

		log.Println(article)
	}
}

func writeArticle(client pb.ArticleServiceClient, in *pb.WriteArticleRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.WriteArticle(ctx, in)

	if err != nil {
		log.Fatalf("Error writing an article: %v", err)
	}

	if res.Value {
		log.Println("Written the article successfully")
	} else {
		log.Println("Cannot write the article")
	}
}

func generateFakeArticle() *pb.WriteArticleRequest {
	return &pb.WriteArticleRequest{
		Title:      faker.Word(),
		Body:       faker.Paragraph(),
		CoverImage: faker.URL(),
	}
}

func main() {
	log.Println("Setting up the article client...")

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithBlock())

	log.Printf("Dialing at %v\n", serverAddr)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewArticleServiceClient(conn)

	articleNums := 3

	for i := 0; i < articleNums; i++ {
		writeArticle(client, generateFakeArticle())
	}

	printArticles(client)
	printArticle(client, "abcd")
}
