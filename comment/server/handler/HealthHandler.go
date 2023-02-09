package handler

import (
	"context"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
)

type HealthHandler struct {
	c.UnsafeHealthServiceServer
}

func (hh *HealthHandler) GetHealth(context.Context, *c.RetrieveHealth) (*c.Health, error) {
	// This application has not external dependencies as ddbb so we needn't check anymore
	return &c.Health{Status: "UP"}, nil
}
