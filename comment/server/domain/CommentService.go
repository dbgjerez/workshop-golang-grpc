package domain

import (
	context "context"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
)

type CommentService struct {
	server c.UnimplementedCommentServiceServer
}

func (service *CommentService) Retrieve(context.Context, *c.RetrieveRequest) (*c.Comments, error) {
	return &c.Comments{}, nil
}
