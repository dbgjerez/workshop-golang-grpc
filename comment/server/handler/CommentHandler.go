package handler

import (
	"context"
	"log"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"github.com/dbgjerez/workshop-golang-grpc/comment/server/domain"
)

type CommentHandler struct {
	c.UnimplementedCommentServiceServer
	comments []*c.Comment
}

func NewServer() *CommentHandler {
	store := domain.NewStore()
	comments, err := store.ReadStore()
	if err != nil {
		log.Fatalf("Error reading the store: %v", err)
	}
	return &CommentHandler{comments: comments}
}

func (s *CommentHandler) Retrieve(ctx context.Context, rq *c.RetrieveRequest) (*c.Comments, error) {
	log.Printf("Request: %s", rq.String())
	return &c.Comments{Comments: s.FilterComments(rq.IdObject, rq.TypeObject)}, nil
}

func (s *CommentHandler) FilterComments(idObject int32, typeOject string) []*c.Comment {
	var res []*c.Comment
	if idObject > 0 || typeOject != "" {
		for _, c := range s.comments {
			if (c.IdObject > 0 && typeOject == "" && c.IdObject == idObject) ||
				(typeOject != "" && idObject <= 0 && c.TypeObject == typeOject) ||
				(c.IdObject == idObject && c.TypeObject == typeOject) {
				res = append(res, c)
			}
		}
	} else {
		res = s.comments
	}
	return res
}
