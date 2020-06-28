package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	"user/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal("Could not dial to server ")
	}
	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CreateUser(ctx, client, uint32(5790))
	GetAllUsers(ctx, client)
	SearchForUser(ctx, client, "hangnhat")
}

func CreateUser(ctx context.Context, client pb.UserServiceClient, id uint32) {
	createUserRequest := pb.CreateUserRequest{
		User: &pb.User{
			UserID:    id,
			UserName:  fmt.Sprintf("hangnhat%d", id),
			UserEmail: fmt.Sprintf("hangnhat%d@gmail.com", id),
			Gender:    pb.User_MALE,
			Age:       30,
		},
	}
	res, err := client.CreateUser(ctx, &createUserRequest)
	if err != nil {
		log.Fatal("Loi cmm", err)
	}
	log.Printf("User created: %s", res.GetMessage())
}
func GetAllUsers(ctx context.Context, client pb.UserServiceClient) {
	stream, _ := client.GetAllUsers(ctx, &pb.Empty{})
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("Shit happened, ", err)
		}
		log.Printf("Getting user with id:%d with name:%s", res.UserID, res.UserName)
	}
}

func SearchForUser(ctx context.Context, client pb.UserServiceClient, keyword string) {
	stream, _ := client.FindUserByName(ctx, &pb.FindUserByNameRequest{Keyword: keyword})

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("Loi cu no roi", err)
		}
		log.Printf("Found user with name %s from keyword %s", response.UserName, keyword)
	}
}
