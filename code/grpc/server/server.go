package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	message "github.com/husterdjx/PS-intership-devops/code/grpc/proto"
)

func handleSendMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	log.Println("receive message:", req.GetSaySomething())
	resp := &message.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}

func main() {
	_ , err := net.Listen("tcp","localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}
	srv := grpc.NewServer()
	message.RegisterMessageSenderService(srv, &message.MessageSenderService{
		Send: handleSendMessage,
	})
}