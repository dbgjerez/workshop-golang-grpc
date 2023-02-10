package handler

import (
	"context"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HealthHandler struct {
	c.UnimplementedHealthServer
}

func (HealthHandler) Check(context.Context, *c.HealthCheckRequest) (*c.HealthCheckResponse, error) {
	return &c.HealthCheckResponse{Status: c.HealthCheckResponse_SERVING}, nil
}

func (HealthHandler) Watch(*c.HealthCheckRequest, c.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
