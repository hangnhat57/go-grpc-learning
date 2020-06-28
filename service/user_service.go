package service

import (
	"context"
	"fmt"
	"log"
	"user/helper"
	"user/model"
	"user/pb"
)

type UserService struct {
	UserDAO model.UserDAO
}

func NewUserService(userDAO model.UserDAO) *UserService {
	return &UserService{userDAO}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := helper.ContextError(ctx)
	if err != nil {
		return nil, err
	}
	err = s.UserDAO.CreateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}
	res := pb.CreateUserResponse{
		Message: fmt.Sprintf("Create user %s successfully", req.User.UserName),
	}
	return &res, nil
}
func (s *UserService) FindUserByName(req *pb.FindUserByNameRequest, stream pb.UserService_FindUserByNameServer) error {
	log.Printf("Search for user name contains %s", req.GetKeyword())
	err := s.UserDAO.FindUser(req, stream)
	if err != nil {
		log.Fatal("Shit happened ", err)
		return err
	}
	return nil
}
func (s *UserService) GetAllUsers(in *pb.Empty, stream pb.UserService_GetAllUsersServer) error {
	log.Printf("Getting all users from data store")
	err := s.UserDAO.GetAllUsers(stream)
	if err != nil {
		log.Fatal("Shit happened ", err)
		return err
	}
	return nil
}
