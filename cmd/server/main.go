package main

import (
	"log"
	"net"
	"os"
	"user/model"
	"user/pb"
	"user/service"

	"user/generator"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	userDao := model.NewInMemUserDaoWithPreDefineData(generator.GenerateUsers(1000))
	userService := service.NewUserService(userDao)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Something error with port 8080 ", err)
	}
	pb.RegisterUserServiceServer(grpcServer, userService)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Something went wrong", err)
		os.Exit(1)
	}
}
