package grpc

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "protos"
	"usrsvc/dao"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUsersInfo(ctx context.Context, user *pb.User) (*pb.User, error) {
	userInfo := dao.UserDao.Get(ctx, user.Id)
	if len(userInfo) == 0 {
		return nil, status.Errorf(codes.NotFound, "user not exits %d", user.Id)
	}
	return &pb.User{
		Name: cast.ToString(userInfo["name"]),
		Id:   cast.ToInt64(userInfo["id"]),
	}, nil
}
