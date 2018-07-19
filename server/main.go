package main

import (
	"log"
	"net"
	pb "github.com/anjuramdas/gRPC_code/list"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type List struct {
	id       int32
	address  string
	replicas []string
}

func (L *List) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Input Message %v", in.Request)
	return &pb.HelloResponse{Reply: "hi client"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	l := List{}
	l.id = 1
	l.address = "localhost:50051"
	s := grpc.NewServer()
	pb.RegisterListServer(s, &l)
	s.Serve(lis)
}

