package grpc

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"protoc"
	"skusvc/dao"
)

type SkuServer struct {
	protoc.UnimplementedHelloServiceServer
}

func (s *SkuServer) DecreaseStock(ctx context.Context, sku *protoc.Sku) (*protoc.Sku, error) {
	// 获取商品信息
	info := dao.SkuDao.Get(ctx, sku.Id)
	if len(info) == 0 {
		return nil, status.Errorf(codes.NotFound, "sku not found")
	}
	// 进行扣减库存
	decrRes, err := dao.SkuDao.Decr(ctx, sku.Id, sku.Num)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	if affected, _ := decrRes.RowsAffected(); affected == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "stock not enough")

	}
	return &protoc.Sku{
		Name:  cast.ToString(info["name"]),
		Id:    cast.ToInt64(info["id"]),
		Price: cast.ToInt32(info["price"]),
		Num:   cast.ToInt32(info["num"]) - sku.Num,
	}, nil
}

func (s *SkuServer) Receive(ctx context.Context, msg *protoc.HelloMsg) (*protoc.HelloMsg, error) {
	// Implement the Receive method
	return &protoc.HelloMsg{Msg: "Hello, " + msg.Msg}, nil
}
