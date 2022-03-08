package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/roguexray007/grpc-demo/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("port", 50051, "the server port")

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("reciever %s", req.GetName())
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcs := grpc.NewServer()
	pb.RegisterGreeterServer(grpcs, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcs.Serve(lis); err !=nil {
		log.Fatalf("failed to serve %v", err)
	}
}