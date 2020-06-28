package model

import (
	"context"
	"errors"
	"log"
	"strings"
	"user/helper"
	"user/pb"
)

var (
	ErrorExisted = errors.New("Already existed!")
)

type UserDAO interface {
	CreateUser(ctx context.Context, user *pb.User) error
	FindUser(req *pb.FindUserByNameRequest, stream pb.UserService_FindUserByNameServer) error
	GetAllUsers(stream pb.UserService_GetAllUsersServer) error
}

type InMemUserDao struct {
	allUsers map[uint32]*pb.User
}

func NewInMemUserDao() *InMemUserDao {
	return &InMemUserDao{make(map[uint32]*pb.User)}
}
func NewInMemUserDaoWithPreDefineData(data map[uint32]*pb.User) *InMemUserDao {
	return &InMemUserDao{allUsers: data}
}
func (userDao *InMemUserDao) CreateUser(ctx context.Context, user *pb.User) error {
	if userDao.allUsers[user.UserID] != nil {
		return ErrorExisted
	}
	userDao.allUsers[user.UserID] = user
	return nil
}

func (userDao *InMemUserDao) FindUser(res *pb.FindUserByNameRequest, stream pb.UserService_FindUserByNameServer) error {
	for _, user := range userDao.allUsers {
		err := helper.ContextError(stream.Context())
		if err != nil {
			return err
		}
		if strings.Contains(user.UserName, res.Keyword) {
			log.Printf("Found user %s", user.UserName)
			err := stream.Send(user)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (userDao *InMemUserDao) GetAllUsers(stream pb.UserService_GetAllUsersServer) error {
	for _, user := range userDao.allUsers {
		err := helper.ContextError(stream.Context())
		if err != nil {
			return err
		}
		err = stream.Send(user)
		if err != nil {
			return err
		}
	}
	return nil
}
