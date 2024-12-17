package grpc

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "protos"
	"skusvc/dao"
)

type SkuServer struct {
	pb.UnimplementedSkuServiceServer
}

func (s *SkuServer) DecreaseStock(ctx context.Context, sku *pb.Sku) (*pb.Sku, error) {
	info := dao.SkuDao.Get(ctx, sku.Id)
	if len(info) == 0 {
		return nil, status.Errorf(codes.NotFound, "sku not exits")
	}
	decrRes, err := dao.SkuDao.Decr(ctx, sku.Id, sku.Num)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	if affected, _ := decrRes.RowsAffected(); affected == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "not enough sku")
	}
	return &pb.Sku{
		Name:  cast.ToString(info["name"]),
		Id:    cast.ToInt64(info["id"]),
		Price: cast.ToInt32(info["price"]),
		Num:   sku.Num,
	}, nil

}
