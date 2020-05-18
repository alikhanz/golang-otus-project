package rotator

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/resources"
	"github.com/alikhanz/golang-otus-project/pkg/pb"
)

type RotatorServer struct {
	validator *Validator
	res *resources.Resources
}

func NewRotatorServer(res *resources.Resources) *RotatorServer {
	v := NewValidator()
	return &RotatorServer{validator: v, res: res}
}

func (r RotatorServer) SelectBanner(
	ctx context.Context,
	request *pb.SelectBannerRequest,
) (*pb.SelectBannerResponse, error) {
	panic("implement me")
}

func (r RotatorServer) AddBannerToSlot(
	ctx context.Context,
	request *pb.AddBannerToSlotRequest,
) (*pb.AddBannerToSlotResponse, error) {
	panic("implement me")
}

func (r RotatorServer) RemoveBannerFromSlot(
	ctx context.Context,
	request *pb.RemoveBannerFromSlotRequest,
) (*pb.RemoveBannerFromSlotResponse, error) {
	panic("implement me")
}
