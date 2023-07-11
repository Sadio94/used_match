package service

import (
	"context"
	"github.com/Sadio94/micro_service/proto"
)

type GoodSvc struct {
	proto.UnimplementedGoodsServer
}

func (this *GoodSvc) GetGoodsByRoom(ctx context.Context, in *proto.GetGoodsByRoomReq) (*proto.GetGoodsByRoomResp, error) {
	out := &proto.GetGoodsByRoomResp{}

	return out, nil
}

func (this *GoodSvc) GetGoodDetail(ctx context.Context, in *proto.GetGoodDetailReq) (*proto.GetGoodDetailResp, error) {
	out := &proto.GetGoodDetailResp{}

	return out, nil
}
