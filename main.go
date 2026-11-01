package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/everestp/blocker/node"
	"github.com/everestp/blocker/proto"
	"google.golang.org/grpc"
)

func main() {
	newNode := node.NewNode()

	// grpc.NewServer() runs insecure by default if no credentials option is passed
	grpcServer := grpc.NewServer()

	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterNodeServer(grpcServer, newNode)
	fmt.Println("node is running on port:", ":3001")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func makeTransaction() {
	// Fixed: Dial port 3001 to match where the server is listening
	client, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer client.Close()

	c := proto.NewNodeClient(client)
	tx := &proto.Transaction{
		Version: 1,
	}
	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatalf("could not handle transaction: %v", err)
	}

	fmt.Println("Transaction sent successfully!")
}
