package handler

import (
	"comment/domain"
	"context"
)

type Server struct {
	domain.UnimplementedCommentServiceServer
}

func (s *Server) Retrieve(context.Context, *domain.RetrieveRequest) (*domain.Comments, error) {
	return &domain.Comments{
		Commets: []*domain.Comment{
			&domain.Comment{
				IdComment:  1,
				IdObject:   12,
				TypeObject: "film",
				IdUser:     20,
				Comment:    "",
			},
		},
	}, nil
}
