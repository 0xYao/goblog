package main

import (
	"0AlexZhong0/goblog/internal/data"
	pb "0AlexZhong0/goblog/internal/generated/api/protobuf/user_service"

	"context"

	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:50051"
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

func deleteUser(c pb.UserServiceClient, userId string) {
	ctx := context.Background()

	_, err := c.DeleteUser(ctx, &pb.DeleteUserRequest{Id: userId})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Deleted the user successfully")
}

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial user service: %v", err)
	}

	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)

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

	for _, item := range createUserRequests {
		deleteUser(userClient, item.Id)
	}

	// expect all the user to be nil
	for _, item := range createUserRequests {
		printUser(userClient, item.Id)
	}
}
