package main

import (
	"io"
	"log"
	"time"

	"0AlexZhong0/goblog/config"
	"0AlexZhong0/goblog/internal/client"
	"0AlexZhong0/goblog/internal/data"
	articlePb "0AlexZhong0/goblog/internal/generated/api/protobuf/article_service"
	"0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func printArticle(client articlePb.ArticleServiceClient, articleId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	article, err := client.GetArticle(ctx, &articlePb.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		log.Fatalf("Failed to get article, %v", err)
	}

	log.Println(article)
}

func printArticles(client articlePb.ArticleServiceClient) {
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

func writeArticle(client articlePb.ArticleServiceClient, in *articlePb.WriteArticleRequest) {
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

func generateFakeArticle() *articlePb.WriteArticleRequest {
	return &articlePb.WriteArticleRequest{
		Title:      faker.Word(),
		Body:       faker.Paragraph(),
		CoverImage: data.RandomImageUrl,
	}
}

func getUsers(userClient user_service.UserServiceClient) []*user_service.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users := make([]*user_service.User, 0)
	stream, err := userClient.GetUsers(ctx, &emptypb.Empty{})

	if err != nil {
		return users
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return users
		}

		users = append(users, user)
	}

	return users
}

func main() {
	config.LoadConfig()
	// userClient, closeUserClientConn, err := client.NewUserClient()

	// defer closeUserClientConn()

	// if err != nil {
	// 	panic(err)
	// }

	articleClient, closeArticleClientConn, err := client.NewArticleClient()
	defer closeArticleClientConn()

	if err != nil {
		panic(err)
	}

	// users := getUsers(userClient)
	// log.Println(users)

	// articleNums := 3

	// for i := 0; i < articleNums; i++ {
	// 	writeArticle(client, generateFakeArticle())
	// }

	printArticles(articleClient)
	printArticle(articleClient, "abcd")
}
