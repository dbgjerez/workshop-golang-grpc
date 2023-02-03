package handler

import (
	"comment/domain"
	"context"
)

var store = domain.NewStore()

type Server struct {
	domain.UnimplementedCommentServiceServer
}

func (s *Server) Retrieve(context.Context, *domain.RetrieveRequest) (*domain.Comments, error) {
	comments, err := store.ReadStore()
	return &domain.Comments{Commets: comments}, err
}
