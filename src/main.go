package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"comment/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	domain.UnimplementedCommentServiceServer
}

func (s *server) Retrieve(context.Context, *domain.RetrieveRequest) (*domain.Comments, error) {
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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	domain.RegisterCommentServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
