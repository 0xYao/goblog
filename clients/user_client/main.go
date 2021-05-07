package main

import (
	"0AlexZhong0/goblog/config"
	"0AlexZhong0/goblog/internal/client"
	"0AlexZhong0/goblog/internal/data"
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"

	"context"

	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func generateFakeCreateUserRequest() *pb.CreateUserRequest {
	return &pb.CreateUserRequest{
		Id:        uuid.NewString(),
		Avatar:    data.RandomImageUrl,
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
	}
}

func printUser(c pb.UserServiceClient, userId string) {
	ctx := context.Background()

	user, err := c.GetUser(ctx, &pb.GetUserRequest{Id: userId})
	if err != nil {
		log.Println(err)
	}

	log.Printf("The user is:\n\n%v\n\n", user)
}

func insertUser(c pb.UserServiceClient, in *pb.CreateUserRequest) {
	ctx := context.Background()

	_, err := c.CreateUser(ctx, in)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Inserted the user successfully")
}

func main() {
	config.LoadConfig()
	userClient, closeUserClientConn, err := client.NewUserClient()

	defer closeUserClientConn()

	if err != nil {
		panic(err)
	}

	// interacting with the server
	userNums := 3
	createUserRequests := []*pb.CreateUserRequest{}

	for i := 0; i < userNums; i++ {
		createUserRequests = append(createUserRequests, generateFakeCreateUserRequest())
	}

	for _, req := range createUserRequests {
		insertUser(userClient, req)
	}

	for _, item := range createUserRequests {
		printUser(userClient, item.Id)
	}
}
