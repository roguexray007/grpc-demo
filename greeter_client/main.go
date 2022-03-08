package main

import (
	"context"
	"flag"
	pb "github.com/roguexray007/grpc-demo/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	default_name = "RTZ"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", default_name, "name to greet")
)

func main(){
	flag.Parse()

	conn, err:= grpc.Dial(*addr, grpc.WithInsecure())
	if err !=nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name,})
	if err != nil {
		log.Fatalf("could not greet %v", err)
	}
	log.Printf("Greetings %s", resp.GetMessage())
}