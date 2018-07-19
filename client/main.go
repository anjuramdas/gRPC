package main

import (
	"log"

	pb "github.com/anjuramdas/gRPC_code/list"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051" //client is connecting with server
	defaultName = "list"
)

func Hi(client pb.ListClient, in *pb.HelloRequest) {
	resp, err := client.Hello(context.Background(), in)
	if err != nil {
		log.Fatalf("node couldn't insert %v", err)
	}

	log.Printf("Reply from server1 %s", resp.Reply)
}
func main() {
	// Setup connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure()) //To connect with server
	if err != nil {
		log.Fatalf("Didn't connect %v", err)
	}
	defer conn.Close() //the defer call to properly close the connection when the function returns
	client := pb.NewListClient(conn)
	in := &pb.HelloRequest{Request: "Hi"}
	Hi(client, in)
}

