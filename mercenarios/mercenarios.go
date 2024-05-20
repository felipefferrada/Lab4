package main

import (
	"context"
	"log"
	"time"

	pb "github.com/felipefferrada/Lab4/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.35.169.67:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := &pb.Message{Text: "Hello, Server!"}
	r, err := c.SendMessage(ctx, message)
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Server response: %s", r.Text)
}
