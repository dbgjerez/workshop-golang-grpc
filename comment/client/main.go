package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	c "github.com/dbgjerez/workshop-golang-grpc/comment/grpc"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	host    = flag.String("host", "localhost", "The server host")
	idObj   = flag.Int("idObj", -1, "The object id")
	typeObj = flag.String("type", "", "The server host")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Call to addr: %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := c.NewCommentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.Retrieve(ctx)
	if err != nil {
		if err := stream.Send(&c.RetrieveRequest{IdObject: int32(*idObj), TypeObject: *typeObj}); err != nil {
			log.Fatalf("send stream: %v", err)
		}
	}
	comments, err := stream.CloseAndRecv()
	log.Println(comments)
}
