package domain

import (
	context "context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type CommentService struct {
	server UnimplementedCommentServiceServer
}

func (service *CommentService) Retrieve(context.Context, *RetrieveRequest) (*Comments, error) {
	return &Comments{}, nil
}

func (service *CommentService) Insert(context.Context, *Comment) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
