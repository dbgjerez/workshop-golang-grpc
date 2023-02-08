package handler

import (
	"context"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"github.com/dbgjerez/workshop-golang-grpc/comment/server/domain"
)

var store = domain.NewStore()

type Server struct {
	c.UnimplementedCommentServiceServer
}

func (s *Server) Retrieve(context.Context, *c.RetrieveRequest) (*c.Comments, error) {
	comments, err := store.ReadStore()
	return &c.Comments{Comments: comments}, err
}
