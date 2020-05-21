package rotator

import (
	"context"
	"errors"

	//"errors"

	"github.com/alikhanz/golang-otus-project/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	validator *Validator
	rotator *Rotator
}

func NewRotatorServer(rotator *Rotator) *Server {
	v := NewValidator(rotator)
	return &Server{validator: v, rotator: rotator}
}

func (r Server) HitBanner(ctx context.Context, request *pb.HitBannerRequest) (*pb.HitBannerResponse, error) {
	err := r.rotator.HitBanner(uint(request.BannerId), uint(request.SlotId), uint(request.SdgId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.HitBannerResponse{}, nil
}

func (r Server) SelectBanner(
	ctx context.Context,
	request *pb.SelectBannerRequest,
) (*pb.SelectBannerResponse, error) {
	banner, err := r.rotator.SelectBanner(uint(request.SlotId), uint(request.SdgId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.SelectBannerResponse{BannerId: uint32(banner.ID)}, nil
}

func (r Server) AddBannerToSlot(
	ctx context.Context,
	request *pb.AddBannerToSlotRequest,
) (*pb.AddBannerToSlotResponse, error) {
	err := r.rotator.AddBannerToSlot(uint(request.BannerId), uint(request.BannerId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.AddBannerToSlotResponse{}, nil
}

func (r Server) RemoveBannerFromSlot(
	ctx context.Context,
	request *pb.RemoveBannerFromSlotRequest,
) (*pb.RemoveBannerFromSlotResponse, error) {
	err := r.rotator.RemoveBannerFromSlot(uint(request.BannerId), uint(request.BannerId))
	err = r.handleErr(err)

	if err != nil {
		return nil, err
	}

	return &pb.RemoveBannerFromSlotResponse{}, nil
}

func (r Server) handleErr(err error) error {
	var errNotFound *NotFoundError

	if err == nil {
		return nil
	}

	if errors.As(err, &errNotFound) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}