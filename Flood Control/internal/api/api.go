package api

import (
	"context"
	"task/internal/service"
	"task/pkg/generated/proto/flood_control"
)

type FloodControlApi struct {
	service service.FloodControlService
	flood_control.UnimplementedFloodControlServiceServer
}

func NewFloodControlApi(service service.FloodControlService) *FloodControlApi {
	return &FloodControlApi{
		service:                                service,
		UnimplementedFloodControlServiceServer: flood_control.UnimplementedFloodControlServiceServer{},
	}
}

func (f *FloodControlApi) Check(ctx context.Context, in *flood_control.CheckRequest) (*flood_control.CheckResponse, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
	}
	allowed, err := f.service.Check(ctx, in.UserId)
	return &flood_control.CheckResponse{Allowed: allowed}, err
}
