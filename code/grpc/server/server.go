package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "PS_intership/code/grpc/proto"
)


const (
	port = ":50053"//在put-forward上的映射端口
)


type server struct {
	pb.UnimplementedMessageSenderServer
}

func (s *server) Send(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", in.GetSaySomething())
	return &pb.MessageResponse{ResponseSomething: "Hello " + in.GetSaySomething()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageSenderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
