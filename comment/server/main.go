package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"github.com/dbgjerez/workshop-golang-grpc/comment/server/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := new(handler.Server)
	c.RegisterCommentServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}