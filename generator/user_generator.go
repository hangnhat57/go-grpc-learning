package generator

import (
	"user/pb"

	"github.com/Pallinder/go-randomdata"
)

func GenerateUsers(max int) map[uint32]*pb.User {
	users := make(map[uint32]*pb.User)
	for i := 0; i < max; i++ {
		users[uint32(i)] = &pb.User{
			UserID:    uint32(i),
			UserName:  randomdata.SillyName(),
			UserEmail: randomdata.Email(),
			Gender:    pb.User_MALE,
			Age:       uint32(randomdata.Number(1, 90)),
		}

	}
	return users
}
